package utils

import "strconv"

func ToUint(s string) (uint, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	return uint(value), err
}

func ToInt(s string) (int, error) {
	value, err := strconv.ParseInt(s, 10, 64)
	return int(value), err
}

func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
