package utils

import (
	"fmt"
	"strconv"
	"time"
)

func Pointer[T any](v T) *T {
	return &v
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StringToBool(param string) (*bool, error) {
	if param == "" {
		return nil, nil
	}
	boolParam, err := strconv.ParseBool(param)
	if err != nil {
		return nil, fmt.Errorf("invalid filter value")
	}
	return &boolParam, nil
}
