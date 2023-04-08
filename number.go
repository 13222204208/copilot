package copilot

import (
	"fmt"
	"time"
)

func UserNum() string {

	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	startTimestamp := time.Now().Unix()
	//获得时间戳
	startTimeStr := time.Unix(startTimestamp, 0).Format("20060102") //把时间戳转换成时间,并格式化为年月日
	nowTime := Sup(hour, 2) + Sup(minute, 2) + Sup(second, 2)
	randNum := GenerateCode()

	code := startTimeStr + nowTime + randNum
	code = code[2:]
	return code
}

// 对长度不足n的数字前面补0
func Sup(i int, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
