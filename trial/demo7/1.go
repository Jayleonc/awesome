package main

import (
	"fmt"
)

type user struct {
	FirstName string
	LastName  string
}

func main() {
	users := make(map[string]user)
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}
	//for key, value := range users {
	//	fmt.Println(key, value)
	//}
	//
	//user1, ok1 := users["Bill"]
	//user2, ok2 := users["Ford"]
	//fmt.Println(user1, ok1)
	//fmt.Println(user2, ok2)

	var result StatNodeErrorRate
	fmt.Println(result)
	fmt.Println(StatNodeErrorRate{}.New(0))

}

func arrayIsEmpty[T []int | []float32 | []user](a T) bool {
	return len(a) != 0
}

type StatNodeErrorRate struct {
	ErrorRate []float64 `json:"errorRate" gorm:"-"`
	Time      []string  `json:"time" gorm:"-"`
}

func (StatNodeErrorRate) New(dataLen int) StatNodeErrorRate {
	result := StatNodeErrorRate{
		ErrorRate: make([]float64, dataLen),
		Time:      make([]string, dataLen),
	}
	return result
}
