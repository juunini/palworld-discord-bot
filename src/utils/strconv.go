package utils

import "strconv"

func ToUint(s string) (uint, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	return uint(value), err
}
