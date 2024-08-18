package helpers

import (
	"strconv"
	"strings"
	"time"
)

func InterfaceToString(v interface{}) string {
	res, ok := v.(string)
	if !ok {
		return ""
	}
	return res
}

func InterfaceToTime(v interface{}, format string) *time.Time {
	res, ok := v.(string)
	if !ok {
		return nil
	}
	dt, err := time.Parse(format, res)
	if err != nil {
		return nil
	}
	return &dt
}

func InterfaceToBoolean(v interface{}) bool {
	res, ok := v.(string)
	if !ok {
		return false
	}

	if strings.ToUpper(res) == "TRUE" {
		return true
	}

	return false
}

func InterfaceToInt64(v interface{}) int64 {
	res, ok := v.(string)
	if !ok {
		return 0
	}
	inCon, err := strconv.Atoi(res)
	if err != nil {
		return 0
	}
	return int64(inCon)
}
