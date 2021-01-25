package NxTime

import (
	"fmt"
	"math"
	"time"
	"errors"
	"strconv"
	"strings"
)

// 获取时间的时间戳
func TimeStamp(tm *time.Time) string {
	return tm.Format("2006-01-02 15:04:05")
}


///////////////////////////////////////////////////////////////////////
/* >>>>>>>>>>>>>>>>>>>>>>>>>> 时间差处理 <<<<<<<<<<<<<<<<<<<<<<<<<<<< */
///////////////////////////////////////////////////////////////////////
///// 通过(Time) Sub() 函数计算时间之差后，对此结果进行分析处理 /////

// 回复格式化的字符串 比如：2h-12m-6s，第二个返回值为正负
func IntervalString(d time.Duration) (string, bool) {
	h, m, s, f := IntervalHMS(d)
	stamp := fmt.Sprintf("%dh-%dm-%ds", h, m, s)
	return stamp, f
}

// 按照时分秒的格式，回复具体的数值
func IntervalHMS(d time.Duration) (h, m, s int, flag bool) {
	duration := int(d)
	if duration >= 0 {
		flag = true
	}

	duration = int(math.Abs(float64(duration)))
	s = duration / 1000000000

	h = s / 3600
	s = s - h * 3600

	m = s / 60
	s = s - m * 60

	return
}



///////////////////////////////////////////////////////////////////////
/* >>>>>>>>>>>>>>>>>>>>>>>>>> 时间的计算 <<<<<<<<<<<<<<<<<<<<<<<<<<<< */
///////////////////////////////////////////////////////////////////////

// 时区转换(默认本时区为东8区, UTC+8)
func TimeTrans(tm *time.Time, timeZone int) time.Time {
	timeDiff := 8 - timeZone //8代表北京时间所处的东8区
	diff := time.Duration(int(time.Hour) * timeDiff)
	return tm.Add(diff)
}

// 时区转换(指定转换到哪个时区)
func TimeTrans2(tm *time.Time, srcZone, destZone int) time.Time {
	timeDiff := destZone - srcZone
	diff := time.Duration(int(time.Hour) * timeDiff)
	return tm.Add(diff)
}



///////////////////////////////////////////////////////////////////////
/* >>>>>>>>>>>>>>>>>>>>>>> 判断周几相关函数 <<<<<<<<<<<<<<<<<<<<<<<<< */
///////////////////////////////////////////////////////////////////////

func IsMonday(tm *time.Time) bool {
	return tm.Weekday() == 1
}

func IsTuesday(tm *time.Time) bool {
	return tm.Weekday() == 2
}

func IsWednesday(tm *time.Time) bool {
	return tm.Weekday() == 3
}

func IsThursay(tm *time.Time) bool {
	return tm.Weekday() == 4
}

func IsFriday(tm *time.Time) bool {
	return tm.Weekday() == 5
}

func IsSaturday(tm *time.Time) bool {
	return tm.Weekday() == 6
}

func IsSunday(tm *time.Time) bool {
	return tm.Weekday() == 0
}

// 返回日期是周几 1~7 分别对应 星期一~星期天
func WhatDay(tm *time.Time) int {
	whatDay := tm.Weekday()
	if whatDay == 0 {
		return 7
	} else {
		return int(whatDay)
	}
}




///////////////////////////////////////////////////////////////////////
/* >>>>>>>>>>>>>>>>>>> 获取时间对象(time.Time) <<<<<<<<<<<<<<<<<<<<<< */
///////////////////////////////////////////////////////////////////////

// 获取一个为"零值"的时间对象
func ZeroTime() time.Time {
	return time.Date(1,time.Month(1), 1, 0, 0, 0, 0, time.UTC)
}

// 根据格式化的时间戳获取
// 年月日和时分秒中间默认是空格
// [sep1]: 年月日之间的分隔符
// [sep2]: 时分秒之间的分隔符
func GetTimeByString(ts, sep1, sep2 string) (time.Time, error) {
	dateTimes := strings.Split(ts, " ")
	if len(dateTimes) != 2 {
		return ZeroTime(), errors.New("Split Error")
	}

	dates := dateTimes[0]
	times := dateTimes[1]

	ymd := strings.Split(dates, sep1)
	hms := strings.Split(times, sep2)
	if len(ymd) != 3 || len(hms) != 3 {
		return ZeroTime(), errors.New("Split Error, sep1 or sep2 wrong")
	}

	year,  err1 := strconv.Atoi(ymd[0])
	month, err2 := strconv.Atoi(ymd[1])
	day,   err3 := strconv.Atoi(ymd[2])
	hour,  err4 := strconv.Atoi(hms[0])
	min,   err5 := strconv.Atoi(hms[0])
	sec,   err6 := strconv.Atoi(hms[0])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		return ZeroTime(), errors.New("Strconv Error")
	}

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local), nil
}