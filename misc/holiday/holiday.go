// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package holiday

import (
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/ysicing/ext/ztime"
)

// Holiday day
type Holiday struct {
	Day       string `json:"day"`
	IsTiaoxiu bool   `json:"is_tiaoxiu"`
	Name      string `json:"name"`
	NeedWork  bool   `json:"need_work"`
}

// Get get
func Get(day string) Holiday {
	var h Holiday
	d := HolidayUtil.GetHoliday(day)
	h.Day = day
	if d == nil {
		t, _ := ztime.TimeParse("2006-01-02", day)
		week := int(t.Weekday())
		if week == 0 || week == 7 || week == 6 {
			h.NeedWork = false
			h.Name = "双休日"
			h.IsTiaoxiu = false
			return h
		}
		h.NeedWork = true
		h.Name = "工作日"
		h.IsTiaoxiu = false
		return h
	}
	h.IsTiaoxiu = d.IsWork()
	h.Name = d.GetName()
	h.NeedWork = false
	return h
}
