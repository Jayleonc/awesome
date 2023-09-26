package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Movie struct {
		Title  string   `json:"title"`
		Year   int      `json:"year"`
		Actors []string `json:"actors"`
	}

	var movie = []Movie{
		{Title: "不能说的秘密", Year: 2007, Actors: []string{"JayChou", "Jayleonc"}},
		{Title: "天台爱情", Year: 2013, Actors: []string{"JayChou", "Ashin"}},
		{Title: "Intouchables", Year: 2011, Actors: []string{"Olivier Nakache", "Eric Toledano"}},
	}

	// struct 转 json
	marshal, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", marshal)

	// json 转 struct
	var title []struct{ Title string }
	err = json.Unmarshal(marshal, &title)
	if err != nil {
		panic(err)
	}

	toMap := JsonToMap(marshal)
	fmt.Printf("%s\n", toMap)

	fmt.Println("actors:")
	for i := range toMap {
		m := toMap[i]
		fmt.Printf("map 第 %d 个：%s\n", i, m)
	}

	beforeMap := []map[string]interface{}{
		{"id": "123", "user_name": "酒窝猪", "address": []map[string]interface{}{{"address": "address01"}, {"address": "address02"}}},
		{"id": "456", "user_name": "酒窝鸡", "address": []map[string]interface{}{{"address": "address01"}, {"address": "address02"}}},
		{"id": "789", "user_name": "酒窝狗", "address": []map[string]interface{}{{"address": "address01"}, {"address": "address02"}}},
	}

	toJson := MapToJson(beforeMap)
	fmt.Printf("%s\n", toJson)
}

func JsonToMap(data []uint8) []map[string]interface{} {
	var mapResult []map[string]interface{}

	err := json.Unmarshal(data, &mapResult)
	if err != nil {
		panic(err)
	}
	return mapResult

}

func MapToJson(data []map[string]interface{}) []uint8 {
	marshal, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	return marshal
}
