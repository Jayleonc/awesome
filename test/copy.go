package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	s := "2023-01-01"
	e := "2023-01-04"
	dates1 := GetBetweenDates1(s, e)
	fmt.Println(dates1)
	sort.SliceStable(dates1, func(i, j int) bool {
		return true
	})
	fmt.Println(dates1[1:])
}

func GetBetweenDates1(startDate, endDate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02"
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		return d
	}
	// 如果结束时间小于开始时间
	if date2.Before(date) {
		return d
	}
	date2Str := fmt.Sprintf("%d年%d月%d日", date2.Year(), int(date2.Month()), date2.Day())
	d = append(d, fmt.Sprintf("%d年%d月%d日", date.Year(), int(date.Month()), date.Day()))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := fmt.Sprintf("%d年%d月%d日", date.Year(), int(date.Month()), date.Day())
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
