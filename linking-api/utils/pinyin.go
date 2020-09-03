/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : pinyin
*/

package utils

import (
	"github.com/Suyghur/pinyin"
	"strings"
	"unicode"
)

func GetLowerPinYinInitials(val string) (result string) {
	pyVal := pinyin.Convert(val)
	for _, r := range pyVal {
		if unicode.IsUpper(r) {
			result = result + string(r)
		}
		if unicode.IsDigit(r) {
			result = result + string(r)
		}
	}
	return strings.ToLower(result)
}

func GetUpperPinYinInitials(val string) (result string) {
	pyVal := pinyin.Convert(val)
	for _, r := range pyVal {
		if unicode.IsUpper(r) {
			result = result + string(r)
		}
	}
	return strings.ToUpper(result)
}

func GetPinYin(val string) string {
	return pinyin.Convert(val)
}
