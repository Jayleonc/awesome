package main

import (
	"fmt"
)

type People struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (p *People) String() string {
	return p.Name
}

func main() {
	//js := `{
	//  "name":"11",
	//"age":12,
	//  "email":"jayleonc@163.com"
	//}`
	p := &People{}
	//err := json.Unmarshal([]byte(js), &p) // 只能解析公共类型的字段，即大写开头的字段
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("people: ", p)
	p.Name = "Jayleonc"
	fmt.Println(p)
}
