package utils

import (
	"time"
)

// GenYesterdayDs 生成昨天的日期戳
func GenYesterdayDs() int {
	now := time.Now().AddDate(0, 0, -1)
	return now.Year()*10000 + int(now.Month())*100 + now.Day()
}

// GenMonthDs 生成base本月初到本月末的日期戳
func GenMonthDs(base time.Time) (int, int) {
	start := base.AddDate(0, 0, -base.Day()+1)
	end := base.AddDate(0, 1, -base.Day())
	return start.Year()*10000 + int(start.Month()*100) + start.Day(),
		end.Year()*10000 + int(end.Month()*100) + end.Day()
}

// GenDate (t-1)00:00:00 - (t-1)23:59:59
func GenDate() (int64, int64) {
	curDay := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	start := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0,
		yesterday.Location())
	end := time.Date(curDay.Year(), curDay.Month(), curDay.Day(), 0, 0, 0, 0,
		curDay.Location())
	return start.Unix(), end.Unix()
}

// GenMs 生成分钟级日期戳
func GenMs(baseTime time.Time) int {
	res := baseTime.Year()*100000000 +
		int(baseTime.Month())*1000000 +
		baseTime.Day()*10000 +
		baseTime.Hour()*100 +
		baseTime.Minute()
	return res
}

// GenDs 生成日期戳
func GenDs(baseTime time.Time) int {
	res := baseTime.Year()*10000 + int(baseTime.Month())*100 + baseTime.Day()
	return res
}
