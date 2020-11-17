// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exfile

import "testing"

func TestRealPath(t *testing.T) {
	type args struct {
		path     string
		addSlash []bool
	}
	tests := []struct {
		name         string
		args         args
		wantRealPath string
	}{
		{
			name: "realpath-ex01",
			args: args{
				path:     "/Users/ysicing/go/src/github.com/ysicing/ext",
				addSlash: []bool{true},
			},
			wantRealPath: "/Users/ysicing/go/src/github.com/ysicing/ext/",
		},
		{
			name: "realpath-ex02",
			args: args{
				path:     "/Users/ysicing/go/src/github.com/ysicing/ext",
				addSlash: []bool{false},
			},
			wantRealPath: "/Users/ysicing/go/src/github.com/ysicing/ext",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRealPath := RealPath(tt.args.path, tt.args.addSlash...); gotRealPath != tt.wantRealPath {
				t.Errorf("RealPath() = %v, want %v", gotRealPath, tt.wantRealPath)
			}
		})
	}
}

func TestWorkDirPath(t *testing.T) {
	type args struct {
		addSlash []bool
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
	}{
		{
			name: "wdpath-ex01",
			args: args{
				addSlash: []bool{true},
			},
		},
		{
			name: "wdpath-ex02",
			args: args{
				addSlash: []bool{false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath := WorkDirPath(tt.args.addSlash...)
			t.Logf("WorkDirPath() = %v, want %v", gotPath, tt.wantPath)
		})
	}
}

func Test_pathAddSlash(t *testing.T) {
	type args struct {
		path     string
		addSlash []bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathAddSlash(tt.args.path, tt.args.addSlash...); got != tt.want {
				t.Errorf("pathAddSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSize2Str(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileSize2Str(tt.args.path); got != tt.want {
				t.Errorf("FileSize2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
