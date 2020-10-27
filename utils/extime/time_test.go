// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package extime

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	years := []int{
		1900, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010,
	}
	for _, year := range years {
		t.Logf("year: %v, 是否闰年: %v", year, IsLeapYear(year))
	}
}

func TestGetMonthDayNum(t *testing.T) {
	months := []map[string]string{
		{"year": "2020", "month": "1"},
		{"year": "2020", "month": "2"},
		{"year": "2020", "month": "3"},
		{"year": "2020", "month": "4"},
		{"year": "2004", "month": "2"},
		{"year": "2003", "month": "2"},
		{"year": "2019", "month": "2"},
	}
	for _, month := range months {
		t.Logf("year: %v, month: %v, days: %v", month["year"], month["month"], GetMonthDayNum(month["year"], month["month"]))
	}
}

func TestGetMonthStartEndUnix(t *testing.T) {
	months := []map[string]string{
		{"year": "2020", "month": "1"},
		{"year": "2020", "month": "2"},
		{"year": "2020", "month": "3"},
		{"year": "2020", "month": "4"},
		{"year": "2004", "month": "2"},
		{"year": "2003", "month": "2"},
		{"year": "2019", "month": "12"},
	}
	for _, month := range months {
		st, et := GetMonthStartEndUnix(month["year"], month["month"])
		t.Logf("year: %v, month: %v, st: %v (%v) et: %v (%v)", month["year"], month["month"], st, UnixInt642String(st), et, UnixInt642String(et))
	}
}

func TestGetMonthAdd(t *testing.T) {
	tests := []int64{0, 1, -1, 10, -10, 12, -12}
	for _, test := range tests {
		t.Logf("value %v res: %v (%v)", test, GetMonthAddStr(test), GetMonthAddInt64(test))
	}
}
