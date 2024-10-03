package utils

import "strconv"

func str2Uint(str string) (uint, error) {
	paramUint64, err := strconv.ParseUint(str, 10, 32)
	paramUint := uint(paramUint64)
	return paramUint, err
}
