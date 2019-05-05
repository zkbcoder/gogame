package util

import "time"

//------------------------------------------------------------------------------
// 时间相关操作
// 获取当前时间戳(单位:s)
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// 获取今天过了多久(单位:s)
func GetTodayPassTime() int64 {
	t := time.Now()
	return int64(t.Hour()*3600 + t.Minute()*60 + t.Second())
}

// 获取每天0点时间戳
func GetTodayBeginTime() int64 {
	tCurTime := time.Now().Unix()
	return tCurTime - GetTodayPassTime()
}

// 获取下一个时刻
func GetNextTimeBySec(uTimeSec int64) int64 {
	if uTimeSec <= GetTodayPassTime() {
		return GetTodayBeginTime() + uTimeSec + 24*3600
	}
	return GetTodayBeginTime() + uTimeSec
}

func GetWeekBeginTime() int64 {
	nWeek := time.Now().Weekday()
	return GetTodayBeginTime() - int64(nWeek)*24*3600 //今天的零点减去周几 比如 今天的周1 就是 减去一天
}

func GetNextWeekTimeBySec(uWeekTimeSec int64) int64 {
	nWeekBeginTime := GetWeekBeginTime()
	nWeekPassTime := time.Now().Unix() - nWeekBeginTime
	if uWeekTimeSec <= nWeekPassTime {
		return nWeekBeginTime + uWeekTimeSec + 24*3600*7
	}
	return nWeekBeginTime + uWeekTimeSec
}
