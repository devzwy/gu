package stringutil

import "strconv"

// 字符串转int 失败时返回-1
func ToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return -1
	}
	return i
}
