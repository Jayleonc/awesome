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

	JsonStr = "{\n    \"non_compliant_projects_count\": \"整治不符合要求的景观照明项目数量（项）\",\n    \"projects_over_50_million\": \"超5000万的景观照明项目数量及总投资额（项、万元）\",\n    \"new_intelligent_posts_count\": \"新建智慧多功能灯杆数量（基）\",\n    \"lighting_digitalization_system\": \"是否建成照明数字化系统\",\n    \"urban_lighting_regulations\": \"是否编制城市照明管理办法\",\n    \"urban_lighting_special_plan\": \"是否编制城市照明专项规划\",\n    \"single_lamp_maintenance_fee\": \"单灯维护费（元）\",\n    \"safety_inspection_frequency\": \"安全隐患定期排查次数（次/平均每月）\",\n    \"deleted\": \"删除标识（0: 正常, 1: 删除）\",\n    \"created_by\": \"创建人\",\n    \"created_time\": \"创建时间\",\n    \"updated_by\": \"更新人\",\n    \"updated_time\": \"更新时间\"\n  }"

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
