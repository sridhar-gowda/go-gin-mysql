package helper

import (
	"os"
	"time"
)

func GetStartHour(timeStamp int64) (start time.Time) {
	prev := timeStamp - (timeStamp % 3600)
	return UnixToTime(prev)
}

func GetStartAndEndHour(timeStamp int64) (start time.Time, end time.Time) {
	prev := timeStamp - (timeStamp % 3600)
	next := prev + 3600
	return UnixToTime(prev), UnixToTime(next)
}
func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func StringToDate(str string) *time.Time {
	time, err := time.Parse("2006-01-02", str)
	if err != nil {
		return nil
	}
	return &time
}

func GetEnv(key string, def string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = def
	}
	return
}
