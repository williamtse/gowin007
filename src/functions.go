package src

import (
	"time"
)

func TimeToUnix(timestr string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("")
	theTime, _ := time.ParseInLocation(timeLayout, timestr+":00", loc)
	sr := theTime.Unix()
	return sr
}
