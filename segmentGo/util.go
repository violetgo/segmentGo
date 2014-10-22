package segmentGo

import (
	"strings"
)

func stringToIntArr(word rune) []int8 {
	temp := make([]int8, 0, 8)
	for cur := word % 10; word/10 != 0 || cur != 0; cur = word % 10 {
		temp = append(temp, int8(cur))
		word = word / 10
	}
	return temp
}

func process(url string) string {
	url = strings.Trim(url, " \t\n\r")
	return url
}

func reverseStr(str []rune) []rune {
	length := len(str)
	result := make([]rune, 0, length)
	for index := length - 1; index >= 0; index-- {
		result = append(result, str[index])
	}
	return result
}
