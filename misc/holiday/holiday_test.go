package holiday

import (
	"reflect"
	"testing"
)

func TestHolidayGet(t *testing.T) {
	type args struct {
		day string
	}
	tests := []struct {
		name string
		args args
		want Holiday
	}{
		{name: "工作日", args: args{day: "2021-04-19"}, want: Holiday{Day: "2021-04-19", IsTiaoxiu: false, NeedWork: true, Name: "工作日"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HolidayGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
