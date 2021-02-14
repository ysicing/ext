// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package ptr

// IntPtr int指针
func IntPtr(p int) *int { return &p }

// Int32Ptr int32指针
func Int32Ptr(p int32) *int32 { return &p }

// Int64Ptr int64指针
func Int64Ptr(p int64) *int64 { return &p }

// StringPtr 字符串指针
func StringPtr(p string) *string { return &p }

// BoolPtr bool指针
func BoolPtr(p bool) *bool { return &p }
