package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
	"time"
)

type Lessee struct {
	Name       string    `json:"name" gorm:"name"`
	ShortName  string    `json:"shortName" gorm:"short_name"`
	Code       string    `json:"code" gorm:"code"`
	CreditCode string    `json:"creditCode" gorm:"credit_code"`
	UnitMan    string    `json:"unitMan" gorm:"unit_man"`
	UnitMobile string    `json:"unitMobile" gorm:"unit_mobile"`
	UnitEmail  string    `json:"unitEmail" gorm:"unit_email"`
	UnitPhone  string    `json:"unitPhone" gorm:"unit_phone"`
	Remark     string    `json:"remark" gorm:"remark"`
	Tenancy    int       `json:"tenancy" gorm:"tenancy"`
	BeginTime  time.Time `json:"beginTime" gorm:"begin_time"`
	EndTime    time.Time `json:"endTime" gorm:"end_time"`
	AppLimit   int       `json:"appLimit" gorm:"app_limit"`
	ManagerId  string    `json:"managerId" gorm:"manager_id"`
	CipherId   string    `json:"cipherId" gorm:"cipher_id"`
	State      int       `json:"state" gorm:"state"`
	BaseModel
}
type BaseModel struct {
	MinBaseModel
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
}

type MinBaseModel struct {
	Id         string    `json:"id" gorm:"column:id;primary_key"`
	CreateBy   string    `json:"createBy" gorm:"column:create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
}

func main() {
	//lessee := Lessee{
	//	Name:       "Jayleonc Company",
	//	ShortName:  "Jay",
	//	CreditCode: "91440101MA9Y9T9K3A",
	//	UnitMan:    "JJ",
	//	UnitMobile: "13333333333",
	//	UnitPhone:  "020-1111111",
	//	UnitEmail:  "jj@gmail.com",
	//	Remark:     "",
	//	Tenancy:    14,
	//	AppLimit:   10,
	//	BeginTime:  time.Now(),
	//	BaseModel: BaseModel{
	//		UpdateTime: time.Now(),
	//		UpdateBy:   "Jayleonc",
	//		MinBaseModel: MinBaseModel{
	//			Id: "08fe01037faeb74ce2e82f92651c2109",
	//		},
	//	},
	//}
	sourceCode := `
		lessee := Lessee{
			Name:       "商密（广州）信息科技有限公司",
			ShortName:  "商密（广州）",
			CreditCode: "91440101MA9Y9T9K3A",
			UnitMan:    "华枫",
			UnitMobile: "13333333333",
			UnitPhone:  "020-1111111",
			UnitEmail:  "huafeng@smsecure.cn",
			Remark:     "",
			Tenancy:    14,
			AppLimit:   10,
			BeginTime:  "",
			BaseModel:  BaseModel{},
			MinBaseModel: MinBaseModel{
				ID: "08fe01037faeb74ce2e82f92651c2109",
			},
		}
	`

	// 移除代码片段中的缩进和换行符
	sourceCode = strings.TrimSpace(sourceCode)
	fields, _ := GetExplicitFields(sourceCode)
	fmt.Println(fields)
}
func convertToJSON(data interface{}) []byte {
	value := reflect.ValueOf(data)
	typeOf := value.Type()

	jsonMap := make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOf.Field(i).Name
		fieldValue := field.Interface()

		// 检查字段的值是否为零值或空值
		if reflect.DeepEqual(fieldValue, reflect.Zero(field.Type()).Interface()) {
			continue
		}

		jsonMap[fieldName] = fieldValue
	}

	jsonData, _ := json.Marshal(jsonMap)
	return jsonData
}

var excludedStructs = []string{"BaseModel", "MinBaseModel"}

func GetInitializedFieldNames(s interface{}) []string {
	var fieldNames []string

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// 检查字段的类型是否为结构体
		if fieldValue.Kind() == reflect.Struct {
			// 递归处理嵌套结构体
			subFieldNames := GetInitializedFieldNames(fieldValue.Interface())
			fieldNames = append(fieldNames, subFieldNames...)
		} else if !fieldValue.IsZero() {
			// 检查字段值是否已初始化
			fieldNames = append(fieldNames, field.Name)
		}
	}

	return fieldNames
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func StructToMap(data interface{}, fields []string) map[string]interface{} {
	result := make(map[string]interface{})
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for _, itemName := range fields {
		field, found := t.FieldByName(itemName)
		if !found {
			continue
		}

		fieldValue := v.FieldByName(itemName)

		// 获取字段名，使用 `json` 标签作为首选，如果没有则使用字段名
		jsonTag := field.Tag.Get("json")
		fieldName := jsonTag
		if fieldName == "" {
			fieldName = field.Name
		}

		// 将字段名转换为小写，并使用下划线分隔单词
		fieldName = toSnakeCase(fieldName)

		// 将字段值添加到结果 map 中
		result[fieldName] = fieldValue.Interface()
	}

	return result
}

func toSnakeCase(name string) string {
	var words []string
	word := ""
	for _, c := range name {
		if c >= 'A' && c <= 'Z' {
			if word != "" {
				words = append(words, strings.ToLower(word))
				word = ""
			}
			word += string(c)
		} else {
			word += string(c)
		}
	}
	if word != "" {
		words = append(words, strings.ToLower(word))
	}
	return strings.Join(words, "_")
}

func GetExplicitFields(sourceCode string) ([]string, error) {
	var fields []string

	// 创建一个文件集合
	fset := token.NewFileSet()

	// 解析源代码片段为表达式
	expr, err := parser.ParseExprFrom(fset, "", sourceCode, 0)
	if err != nil {
		return nil, fmt.Errorf("解析源代码失败: %v", err)
	}

	// 检查解析结果是否为复合字面量
	if compLit, ok := expr.(*ast.CompositeLit); ok {
		// 检查复合字面量的类型是否为结构体
		if structType, ok := compLit.Type.(*ast.StructType); ok {
			// 遍历复合字面量的元素
			for _, elem := range compLit.Elts {
				// 检查元素是否为键值对表达式
				if kv, ok := elem.(*ast.KeyValueExpr); ok {
					// 检查键是否为标识符
					if ident, ok := kv.Key.(*ast.Ident); ok {
						// 检查键对应的字段是否为结构体字段
						for _, field := range structType.Fields.List {
							if field.Names != nil {
								// 遍历结构体字段，比较字段名和键名
								for _, fieldName := range field.Names {
									if fieldName.Name == ident.Name {
										fields = append(fields, ident.Name)
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return fields, nil
}
