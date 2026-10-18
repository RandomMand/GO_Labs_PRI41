package stringutil

import "strings"

func ReverseStr(stroka string) string {
	array := strings.Split(stroka, "")
	var result string
	for i := len(array) - 1; i >= 0; i-- {
		result += array[i]
	}
	return result
}
