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
