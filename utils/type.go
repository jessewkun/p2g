package utils

import (
	"strconv"
)

func ToString(arg interface{}) string {
	switch val := arg.(type) {
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'g', 1, 32)
	case float64:
		return strconv.FormatFloat(val, 'g', 1, 64)
	case string:
		return val
	default:
		return ""
	}
}
