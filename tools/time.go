package tools

import (
	"time"
)

//func Time() int64 {
//	return time.Now().Unix() - 5*60*60
//}

func Time() int64 {
	return time.Now().Unix()
}

func CurrentDate() string {
	timeLayout := "2006-01-02"
	dateTimeStr := time.Unix(Time(), 0).Format(timeLayout)
	return dateTimeStr
}

func CurrentDateTime() string {
	timeLayout := "2006-01-02 15:04:05"
	dateTimeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	return dateTimeStr
}

func MilliSecondTime() int64 {
	return int64(time.Now().UnixNano() / 1e6)
}

func NewYorkTime() int64 {
	format := "2006-01-02 15:04:05"
	day := time.Now().Format("2006-01-02 15:04:05")
	local, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation(format, day, local)
	return t.UnixNano() / 1e9
	//t.UnixNano() / 1e6 毫秒
}

func GetDateTime(date string) (int64, int64) {
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//toDayStartTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	//toDayEndTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Local).Unix()
	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}

func GetTimeArr(start, end string) int64 {
	timeLayout := "2006/01/02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

func StrToDateTime(t string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ymd, _ := time.ParseInLocation(timeLayout, t, loc)
	return ymd
}

func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}
