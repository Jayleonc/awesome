package main

import (
	"fmt"
	"strings"
)

type Field struct {
	Before string // 改变前的值
	After  string // 改变后的值
}

func main() {
	filedMap := make(map[string]Field)
	filedMap["组织关系"] = Field{Before: "ok", After: "okk"}
	filedMap["排序"] = Field{Before: "123", After: "1234"}
	filedMap["备注"] = Field{Before: "你干嘛", After: "你干嘛嗨嗨~"}

	updDesc := "修改"
	objDesc := "组织"
	ids := []string{"23142345262526nk25gw2"}

	res := fmt.Sprintf("【%s】%s, ID: %v", updDesc, objDesc, ids)
	fieldSList := make([]string, 0)
	if len(filedMap) > 0 {
		for k, v := range filedMap {
			fieldSList = append(fieldSList, fmt.Sprintf("%s: 改前\"%s\", 改后\"%s\"", k, v.Before, v.After))
		}
	}
	if len(fieldSList) > 0 {
		str := strings.Join(fieldSList, "; ")
		res = fmt.Sprintf("%s; 修改字段: [ %s ]", res, str)
	}

	fmt.Println(res)
}
