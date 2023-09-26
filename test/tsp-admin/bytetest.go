package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	fmt.Println(humanize.Bytes(3301))
	fmt.Println(3301 / 1024)
	num2, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(3150)/float64(1024)), 64)
	fmt.Println(num2)

	//httpRate, _ := decimal.NewFromFloat(10 / 10).Round(2).Float64()
	////socketRate, _ := decimal.NewFromFloat(0 / 0).Round(2).Float64()
	//soapRate, _ := decimal.NewFromFloat(10 / 0).Round(2).Float64()
	//fmt.Println(httpRate)
	////fmt.Println(socketRate)
	//fmt.Println(soapRate)
	//mailRate, _ := decimal.NewFromFloat(data[i].RequestFailMail / data[i].RequestMail).Round(2).Float64()
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, -1)                //年，月，日   获取一天前的时间
	resTime := getTime.Format("2006-01-02 15:04:05+08") //获取的时间的格式
	fmt.Println(resTime)
	var errArr = make([]int, 9)
	for i := range errArr {
		errArr[i] = i
	}
	for i := 0; i < len(errArr); i++ {
		fmt.Println(errArr[i])
	}

	gin.Default()
}
