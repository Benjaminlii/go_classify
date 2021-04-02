package util

import "strconv"

// StringToUInt string类型转换为uint类型，十进制
func StringToUInt(str string) (uint, error) {
	gotUint64, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(gotUint64), nil
}
