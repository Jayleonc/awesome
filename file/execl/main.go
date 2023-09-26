package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	A = "A"
	B = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
	H = "H"
)

type ColRaw struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
	E string `json:"E"`
	F string `json:"F"`
	G string `json:"G"`
	H string `json:"H"`
}

var JsonStr string
var FileName string

func main() {
	//fmt.Println("请输入 Json 数据：")
	//fmt.Scan(&JsonStr)

	JsonStr = "  {\n  \"ylgy\": \"游乐公园（个）\",\n  \"ylgymj\": \"游乐公园面积（公顷）\",\n  \"kdgy\": \"口袋公园（个）\",\n  \"kdgymj\": \"口袋公园面积（公顷）\",\n  \"path\": \"路径\",\n  \"pathName\": \"路径层级名称\",\n  \"gymj\": \"公园面积\"\n}"

	fmt.Println("请输入文件名称：")
	fmt.Scan(&FileName)

	// 解析 JSON 数据
	data := make(map[string]string)
	if err := json.Unmarshal([]byte(JsonStr), &data); err != nil {
		fmt.Println("JSON 解析错误:", err)
		return
	}

	// 获取键的切片并排序
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	//sort.Strings(keys)

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 创建一个新的工作表
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)
	f.SetDefaultFont("宋体")

	// 将键值对按顺序写入 Excel 表格
	row := 1
	for _, key := range keys {
		value := data[key]
		// 将键存储到第一行，值存储到第二行
		f.SetCellStr(sheetName, fmt.Sprintf("%c%d", toChar(64+row), 1), value)
		f.SetCellStr(sheetName, fmt.Sprintf("%c%d", toChar(64+row), 2), "{."+key+"}")
		row++
	}

	// 设置活动工作表
	f.SetActiveSheet(index)

	// 保存 Excel 文件
	if err := f.SaveAs(FileName + ".xlsx"); err != nil {
		fmt.Println("保存 Excel 文件错误:", err)
		return
	}

	fmt.Println("Excel 文件已生成成功。")
}

func toChar(i int) rune {
	return rune(i)
}
