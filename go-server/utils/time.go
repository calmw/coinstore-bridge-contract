package utils

import (
	"time"
)

func DatetimeToUnix(timeStr string) (int64, error) {
	// 解析特定时间
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err
	}
	return parsedTime.Unix(), nil
}

// TimestampToDatetime 将时间戳转换为北京时间
func TimestampToDatetime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ""
	}
	bjTime := t.In(loc)
	return bjTime.Format("2006-01-02 15:04:05")
}
