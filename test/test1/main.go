package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func main() {
	source := "root:jayleonc@tcp(localhost:3306)/tsp?charset=utf8&parseTime=true&loc=Local&timeout=1000ms"
	var schemaURL string
	mysqlIndex0 := strings.LastIndex(source, "/") + 1
	questionMarkIndex0 := strings.Index(source, "?")
	if mysqlIndex0 > 0 && questionMarkIndex0 > mysqlIndex0 {
		schemaURL = source[:mysqlIndex0] + "information_schema" + source[questionMarkIndex0:]
	} else {
		// 处理在 `global.Source` 中未找到 `/` 或 `?` 或它们出现的顺序不正确的情况。
		log.Fatal("无效的数据库URL")
	}
	fmt.Println(schemaURL)
	db0, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}

	log.Print("检查数据库是否存在...")
	res, err := db0.Query("select * from SCHEMATA where SCHEMA_NAME = 'tsp'")
	if err != nil {
		panic(err)
	}
	fmt.Print(res.Next())
}
