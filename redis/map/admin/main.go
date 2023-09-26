package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	MinBaseModel
	UpdateBy   string    `gorm:"column:update_by"`   // 修改人ID
	UpdateTime time.Time `gorm:"column:update_time"` // 修改时间 格式"2001-01-01 01:01:01.111"
}

type MinBaseModel struct {
	Id         string    `gorm:"column:id;primary_key"` // 主键 32位全小写GUID字符串
	CreateBy   string    `gorm:"column:create_by"`      //创建人ID
	CreateTime time.Time `gorm:"column:create_time"`    // 创建时间 格式"2001-01-01 01:01:01.111"
}
type GatewayObject struct {
	AppId      string            `json:"appId" gorm:"column:app_id" `      // 应用ID
	LesseeId   string            `json:"lesseeId" gorm:"column:lessee_id"` // 机构ID
	KeyId      string            `json:"keyId" gorm:"column:key_id"`       // 密钥ID
	KeyStorage map[string]string `json:"keyStorage" gorm:"-"`              // 密钥存储信息
	Secret     string            `json:"secret" gorm:"secret"`

	//NodeId   string `json:"nodeId" gorm:"column:node_id"`     // 节点ID
	//KeyIndex string `json:"keyIndex" gorm:"column:key_index"` // 密钥索引

}

type CipherStorage struct {
	BaseModel
	CipherId       string    `json:"cipherId" gorm:"column:cipher_id" `              // 签戳密钥ID
	NodeId         string    `json:"nodeId" gorm:"column:node_id" `                  // 时间戳服务器id
	KeyIndex       string    `json:"keyIndex" gorm:"column:key_index" `              // 密钥索引
	First          int8      `json:"first" gorm:"column:first" `                     // 是否密钥生成节点
	SyncCipher     int8      `json:"syncCipher" gorm:"column:sync_cipher" `          // 是否同步签戳密钥
	SyncCert       int8      `json:"syncCert" gorm:"column:sync_cert" `              // 是否同步签戳证书
	SyncCipherTime time.Time `json:"syncCipherTime" gorm:"column:sync_cipher_time" ` // 成功同步签戳密钥时间
	SyncCertTime   time.Time `json:"syncCertTime" gorm:"column:sync_cert_time" `     // 成功同步签戳证书时间
	State          int8      `json:"state" gorm:"column:state" `                     // 状态
}

type Need struct {
	AppId     string            `json:"appId"`
	KeyId     string            `json:"keyId"`
	LesseeId  string            `json:"lesseeId"`
	Storages  map[string]string `json:"storages"`
	NodeState bool              `json:"nodeState"`
}

func main() {
	defer timeTrack(time.Now(), "My Test")
	var err error
	err = db(err)
	needMap, err := GetGatewayObjectMapTest()
	if err != nil {
		return
	}

	// 发送到 Redis 中
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.140:6379",
		Password: "smsecure123",
		DB:       0,
	})

	for i, v := range needMap {
		fmt.Printf("index: %v, value: %v\n", i, v)
		result, err := client.HSetNX(context.Background(), "hash-key-test-2", i, v).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

}

func GetGatewayObjectMap() (map[string]interface{}, error) {
	var objs []GatewayObject
	// 定义参数
	db := DB.Table("tsp_object tob").
		Joins("LEFT JOIN tsp_lessee tl ON tl.id = tob.lessee_id")
	if err := db.Where("tob.state = ?", 0).
		Select("tob.app_id, tob.lessee_id, tl.cipher_id AS key_id").
		Find(&objs).Error; err != nil {
		return nil, err
	}

	needMap := make(map[string]interface{})
	// 计算密钥存储信息
	if objs != nil {
		for i := range objs {
			// 根据keyId获取对应的密钥存储信息
			storages, count, err := GetCipherStorage(objs[i].KeyId)
			if err != nil {
				return nil, err
			}

			if count > 0 {
				var storageMap = make(map[string]string)
				for k := range storages {
					storageMap[storages[k].NodeId] = storages[k].KeyIndex // key = nodeId， value = keyIndex
				}
				objs[i].KeyStorage = storageMap
			}

			var need Need
			need.AppId = objs[i].AppId
			need.KeyId = objs[i].KeyId
			need.LesseeId = objs[i].LesseeId
			need.Storages = objs[i].KeyStorage
			data, err := need.MarshalBinary()
			if err != nil {
				return nil, err
			}
			needMap[need.AppId] = data
		}
	}

	return needMap, nil

}
func GetCipherStorage(cipherId string) ([]CipherStorage, int64, error) {
	// 定义参数
	var storage []CipherStorage
	var count int64
	db := DB.Table("tsp_cipher_storage tcs").
		Joins("LEFT JOIN tsp_node tn ON tn.id = tcs.node_id")
	// 查询
	if err := db.Where("tcs.cipher_id = ? AND tcs.state <> ? AND tn.state = ?",
		cipherId, 9, 0).
		Count(&count).Select("tcs.id, tcs.cipher_id, tcs.node_id, tcs.key_index, tcs.first, tcs.sync_cipher, " +
		"tcs.sync_cert, tcs.sync_cipher_time, tcs.sync_cert_time").
		Find(&storage).Error; err != nil {
		return []CipherStorage{}, 0, err
	}
	return storage, count, nil
}

func (n *Need) MarshalBinary() ([]byte, error) {
	return json.Marshal(n)
}
func db(err error) error {
	dsn := "root:smsecure@140server@tcp(119.29.152.158:13307)/tsp?charset=utf8&parseTime=true&loc=Local&timeout=1000ms"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return err
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
