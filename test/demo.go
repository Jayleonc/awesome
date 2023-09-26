package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func main() {
	//dates := GetBetweenDates("2022-12-14", "2022-12-21")
	//fmt.Println(dates)
	now := time.Now()
	oldLayout := "2006-01-02"
	nowStr := now.Format(oldLayout)
	hourZeroMin := nowStr
	layout := "2006-01-02"
	endTime, _ := time.Parse(layout, hourZeroMin)
	p1, _ := time.ParseDuration("-24h")
	endTime = endTime.Add(p1)
	fmt.Println(endTime.Format("2006-01-02"))
	//ip, err := getClientIp()
	//if err != nil {
	//	return
	//}
	//fmt.Println(ip)
	t1 := time.Date(2023, 1, 6, 10, 0, 0, 100, time.Local)
	//t2 := time.Date(2023, 1, 3, 8, 02, 22, 100, time.Local)
	//sub := timeSub(t1, t2)
	//fmt.Println("相差小时数：", sub)

	if TimeGreaterThanToday(t1) {
		fmt.Println("大于今天")
	}
	if TimeGreaterThanToday2(t1) {
		fmt.Println("大于今天")
	}

	if TimeEqualToday(t1) {
		fmt.Println("等于今天")
	}
}

func GetBetweenDates(sdate, edate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	//date2Str := date2.Format(timeFormatTpl)
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
	return d[:len(d)-1]
}
func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")

}

func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), 0, 0, 0, time.Local)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	sub := t1.Sub(t2)
	fmt.Println("sub:", sub)
	a := int(t1.Sub(t2).Hours())
	return a
}
func TimeEqualToday(startAt time.Time) bool {
	t1, _ := time.Parse("2006-01-02", startAt.Format("2006-01-02"))
	t2, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return t1.Equal(t2)
}

func TimeGreaterThanToday(startAt time.Time) bool {
	return startAt.Format("2006-01-02") > time.Now().Format("2006-01-02")
}

func TimeGreaterThanToday2(startAt time.Time) bool {
	t1, _ := time.Parse("2006-01-02", startAt.Format("2006-01-02"))
	t2, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return t1.After(t2)
}
