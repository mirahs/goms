package time

import "time"


func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}


// UnixTime 1970年01月01日00时00分00秒起至现在的秒数
func UnixTime() int64 {
	return time.Now().Unix()
}

func DateTime2UnixTime(dateTimeStr string) int64 {
	format := "2006-01-02 15:04:05"

	res, _ := time.ParseInLocation(format, dateTimeStr, time.Local)
	return res.Unix()
}

func UnixTime2DateTime(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

func UnixTime2Date(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02")
}

func TodayBeginAdnEndTime() (timeBegin, timeEnd int64) {
	return DayBeginAdnEndTime(UnixTime2Date(UnixTime()))
}

func DayBeginAdnEndTime(dateStr string) (timeBegin, timeEnd int64) {
	timeBegin = DateTime2UnixTime(dateStr + " 00:00:00")
	timeEnd = DateTime2UnixTime(dateStr + " 23:59:59")

	return
}
