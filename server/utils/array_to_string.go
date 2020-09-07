/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : array2string
*/

package utils

import (
	"fmt"
	"strings"
)

func Array2String(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
