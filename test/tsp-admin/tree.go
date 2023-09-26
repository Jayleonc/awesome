package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Node struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Children []*Node `json:"children"`
}
type GroupDetail struct {
	Id       string         `json:"id"`
	ParentId string         `json:"parentId" gorm:"column:parent_id"`
	Name     string         `json:"name" gorm:"column:name" ` // 组织名称
	Child    []*GroupDetail `json:"child" gorm:"-"`
}

func getTree(list []*GroupDetail, parentId string) []*GroupDetail {
	res := make([]*GroupDetail, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			v.Child = getTree(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}

func getTreeRecursive(list []*Node, parentId int) []*Node {
	res := make([]*Node, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			v.Children = getTreeRecursive(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}

func main() {
	//list := []*Node{
	//	{4, 3, "ABA", nil},
	//	{3, 1, "AB", nil},
	//	{1, 0, "A", nil},
	//	{2, 1, "AA", nil},
	//}
	//res := getTreeRecursive(list, 0)
	//bytes, _ := json.MarshalIndent(res, "", "    ")
	//fmt.Printf("%s\n", bytes)

	l := []*GroupDetail{
		{"3d60d39be0a6cf8732e12539724490f4", "0", "ok", nil},
		{"9a0342418fec6e6f6ba3482fd65c7497", "3d60d39be0a6cf8732e12539724490f4", "ok", nil},
		{"06223aa5bf23247ee52705b8c84f9c8f", "3d60d39be0a6cf8732e12539724490f4", "ok", nil},
		{"b9e6056226218ddc112bd8903d62e9cb", "3d60d39be0a6cf8732e12539724490f4", "ok", nil},
		{"bc4ef49661112271f6b0cd955b05b8b5", "b9e6056226218ddc112bd8903d62e9cb", "ok", nil},
		{"cb9f6b3c545a2b59f319d79c4598c16c", "b9e6056226218ddc112bd8903d62e9cb", "ok", nil},
		{"182e22a1b6228dd5499abad57719bbcb", "b9e6056226218ddc112bd8903d62e9cb", "ok", nil},
	}
	tree := getTree(l, "0")
	bytes, _ := json.MarshalIndent(tree, "", "    ")
	fmt.Printf("%s\n", bytes)

	i := 0
	fmt.Println(strconv.Itoa(i))
}
