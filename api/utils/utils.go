package utils

import "time"

func Pointer[T any](v T) *T {
	return &v
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
