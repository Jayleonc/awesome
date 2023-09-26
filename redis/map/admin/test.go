package main

import "sync"

func GetGatewayObjectMapTest() (map[string]interface{}, error) {
	var objs []GatewayObject
	// 定义参数
	db := DB.Table("tsp_object tob").
		Joins("LEFT JOIN tsp_lessee tl ON tl.id = tob.lessee_id")
	if err := db.Where("tob.state = ?", 0).
		Select("tob.app_id, tob.lessee_id, tl.cipher_id AS key_id").
		Find(&objs).Error; err != nil {
		return nil, err
	}

	return processObjects(objs)
}

func processObjects(objs []GatewayObject) (map[string]interface{}, error) {
	// 并发计算密钥存储信息和序列化
	type result struct {
		key   string
		value []byte
		err   error
	}

	ch := make(chan result, len(objs))
	var wg sync.WaitGroup
	wg.Add(len(objs))
	for i := range objs {
		go func(obj GatewayObject) {
			defer wg.Done()
			// 根据 keyId 获取对应的密钥存储信息
			storages, count, err := GetCipherStorage(obj.KeyId)
			if err != nil {
				ch <- result{err: err}
				return
			}
			if count > 0 {
				var storageMap = make(map[string]string, len(storages))
				for k := range storages {
					storageMap[storages[k].NodeId] = storages[k].KeyIndex
				}
				obj.KeyStorage = storageMap
			}
			var need Need
			need.AppId = obj.AppId
			need.KeyId = obj.KeyId
			need.LesseeId = obj.LesseeId
			need.Storages = obj.KeyStorage
			data, err := need.MarshalBinary()
			if err != nil {
				ch <- result{err: err}
				return
			}
			ch <- result{key: obj.AppId, value: data}
		}(objs[i])
	}

	// 等待所有协程完成
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 从通道中读取结果并组成 map 返回
	needMap := make(map[string]interface{})
	for res := range ch {
		if res.err != nil {
			return nil, res.err
		}
		needMap[res.key] = res.value
	}
	return needMap, nil
}
