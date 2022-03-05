package utils

import "time"

const FORMAT = "2006-01-02 15:04:05"

func IsToday(d time.Time) bool {
	now := time.Now()
	return d.Year() == now.Year() && d.Month() == now.Month() && d.Day() == now.Day()
}

// 计算过期时间是否超过l
func WithinLimit(s int64, l int64) bool {
	e := time.Now().Unix()
	// println(e - s)
	return e-s < l
}
