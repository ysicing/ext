// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package ztime

import (
	"fmt"
	"github.com/ysicing/ext/exstr"
	"strings"
	"time"
)

// NowUnixString 当前时间时间戳
func NowUnixString() string {
	return exstr.Int642Str(time.Now().Unix())
}

// NowUnix 当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// TimeParse time parse
func TimeParse(layout, t string) (time.Time, error) {
	return time.Parse(layout, t)
}

// TimeLayoutStr time layout to str
func TimeLayoutStr(layout string) string {
	return time.Now().Format(layout)
}

// NowStr 当前时间
func NowStr() string {
	return TimeLayoutStr("2006-01-02 15:04:05")
}

// NowMin 获取当前分钟
func NowMin() string {
	return TimeLayoutStr("200601021504")
}

// NowHour 获取当前小时
func NowHour() string {
	return TimeLayoutStr("2006010215")
}

// NowDay 获取当前日期
func NowDay() string {
	return TimeLayoutStr("20060102")
}

// NowMonth 获取当前月份
func NowMonth() string {
	return TimeLayoutStr("200601")
}

// NowYear 获取当前年份
func NowYear() string {
	return TimeLayoutStr("2006")
}

// NowShortHour 获取今天具体小时
func NowShortHour() string {
	return strings.ReplaceAll(NowHour(), NowDay(), "")
}

// NowShortMonth 获取当前月份
func NowShortMonth() string {
	return strings.ReplaceAll(NowMonth(), NowYear(), "")
}

// Today0hourUnix 今天0时时间戳
func Today0hourUnix() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
	return t1.Unix()
}

// BeforeNowUnix 历史时间戳
func BeforeNowUnix(old string) (oldunix int64) {
	return time.Now().Unix() - exstr.Str2Int64(old)
}

// UnixInt642String unix转化为字符串
func UnixInt642String(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// UnixString2String unix转化为字符串
func UnixString2String(t string) string {
	return time.Unix(exstr.Str2Int64(t), 0).Format("2006-01-02 15:04:05")
}

// UnixNanoInt642String unix转化为字符串
func UnixNanoInt642String(t int64) string {
	return time.Unix(0, t).Format("2006-01-02 15:04:05")
}

// UnixNanoString2String unix转化为字符串
func UnixNanoString2String(t string) string {
	return time.Unix(0, exstr.Str2Int64(t)).Format("2006-01-02 15:04:05")
}

// GetWeekFristDayUnix 时间
func GetWeekFristDayUnix() int64 {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart.Unix()
}

// GetWeekLastDayUnix 时间
func GetWeekLastDayUnix() int64 {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekNextStart.Unix() - 1
}

// GetWeekDayUnix 时间
func GetWeekDayUnix() (int64, int64) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekStart.Unix(), weekNextStart.Unix() - 1
}

// GetWeekDayUnixString 时间
func GetWeekDayUnixString() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekStart.Format("2006-01-02 15:03:04"), weekNextStart.Format("2006-01-02 15:03:04")
}

// NowAddUnix2Int64 当前时间戳 add
func NowAddUnix2Int64(key string, value time.Duration) int64 {
	lkey := strings.ToLower(key)
	if lkey == "m" || lkey == "minute" {
		return time.Now().Add(time.Minute * value).Unix()
	}
	if lkey == "d" || lkey == "day" {
		return time.Now().Add(time.Hour * 24 * value).Unix()
	}
	if lkey == "w" || lkey == "week" {
		return time.Now().Add(time.Hour * 24 * 7 * value).Unix()
	}
	return time.Now().Add(time.Hour * value).Unix()
}

// NowAddUnix2Str 当前时间戳 add
func NowAddUnix2Str(key string, value time.Duration) string {
	lkey := strings.ToLower(key)
	if lkey == "m" || lkey == "minute" {
		return time.Now().Add(time.Minute * value).String()
	}
	if lkey == "d" || lkey == "day" {
		return time.Now().Add(time.Hour * 24 * value).String()
	}
	if lkey == "w" || lkey == "week" {
		return time.Now().Add(time.Hour * 24 * 7 * value).String()
	}
	return time.Now().Add(time.Hour * value).String()
}

// GetMonthDayNum 获取任已一年月的天数
func GetMonthDayNum(year, month string) int {
	switch month {
	case "01", "03", "05", "07", "08", "10", "12":
		return 31
	case "04", "06", "09", "11":
		return 30
	default:
		if IsLeapYear(exstr.Str2Int(year)) {
			return 29
		}
		return 28
	}
}

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool { //y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

// TimeToUninx 时间转时间戳
func TimeToUninx(layout, t string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation(layout, t, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.Unix()
}

// GetMonthStartEndUnix 某月开始和结束时间戳
func GetMonthStartEndUnix(year, month string) (int64, int64) {
	if exstr.Str2Int(month) < 10 {
		month = fmt.Sprintf("0%v", month)
	}
	st := fmt.Sprintf("%v-%v-01 00:00:00", year, month)
	et := fmt.Sprintf("%v-%v-%v 23:59:59", year, month, GetMonthDayNum(year, month))
	return TimeToUninx("2006-01-02 15:04:05", st), TimeToUninx("2006-01-02 15:04:05", et)
}

// GetMonthAddInt64 前几个月或者后几个月
func GetMonthAddInt64(value int64) int64 {
	year := time.Now().Year()
	mon := exstr.Str2Int64(NowShortMonth())
	// value -1
	if value >= 0 {
		if value+mon > 12 {
			year = year + 1
			value = value + mon - 12
			_, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
			return et
		}
		value = value + mon
		_, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return et
	}
	// value -10 12月
	if value+mon <= 0 {
		year = year - 1
		value = 12 + (value + mon)
		st, _ := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return st
	}
	value = value + mon
	st, _ := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
	return st
}

// GetMonthAddStr 前几个月或者后几个月
func GetMonthAddStr(value int64) string {
	year := time.Now().Year()
	mon := exstr.Str2Int64(NowShortMonth())
	// value -1
	if value >= 0 {
		if value+mon > 12 {
			year = year + 1
			value = value + mon - 12
			_, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
			return UnixInt642String(et)
		}
		value = value + mon
		_, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return UnixInt642String(et)

	}
	// value -10 12月
	if value+mon <= 0 {
		year = year - 1
		value = 12 + (value + mon)
		st, _ := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return UnixInt642String(st)
	}
	value = value + mon
	st, _ := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
	return UnixInt642String(st)
}

// NowMonthAddNum 某个月
func NowMonthAddNum(value int64) (int64, int64) {
	year := time.Now().Year()
	mon := exstr.Str2Int64(NowShortMonth())
	// value -1
	if value >= 0 {
		if value+mon > 12 {
			year = year + 1
			value = value + mon - 12
			st, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
			return st, et
		}
		value = value + mon
		st, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return st, et
	}
	// value -10 12月
	if value+mon <= 0 {
		year = year - 1
		value = 12 + (value + mon)
		st, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
		return st, et
	}
	value = value + mon
	st, et := GetMonthStartEndUnix(exstr.Int642Str(int64(year)), exstr.Int642Str(value))
	return st, et
}
