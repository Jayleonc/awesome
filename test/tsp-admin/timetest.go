package main

import (
	"fmt"
	"time"
)

func main() {
	agoTime, s := GetOneDayAgoTime()
	hours := GetBetweenHours(agoTime, s)
	fmt.Println(hours)
}

func GetBetweenHours(startDate, endDate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
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
	// 如果前一天时间（endDate）大于今天时间（startDate）
	if date2.After(date) {
		return d
	}
	dateStr := date.Format("2006-01-02 15") + ":00"
	d = append(d, date2.Format("2006-01-02 15")+":00")
	for {
		p1, _ := time.ParseDuration("1h")
		date2 = date2.Add(p1)
		date2Str := date2.Format("2006-01-02 15") + ":00"
		d = append(d, date2Str)
		if date2Str == dateStr {
			break
		}
	}
	return d
}

func GetOneDayAgoTime() (string, string) {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, -1)                 //年，月，日   获取一天前的时间
	last24hTime := getTime.Format("2006-01-02 15:04:05") //获取的时间的格式
	nowTimeStr := nowTime.Format("2006-01-02 15:04:05")
	return nowTimeStr, last24hTime
}
