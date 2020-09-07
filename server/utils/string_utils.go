/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : string_utils
*/

package utils

import (
	"math/rand"
	"time"
)

const (
	AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLetters   = AsciiLowercase + AsciiUppercase
	Digits         = "0123456789"
)

// 随机生成指定位数的大写字母和数字的组合
func GetRandomDigitsUpperStr(l int) string {
	str := Digits + AsciiUppercase
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成指定位数的大写字母和数字的组合
func GetRandomDigitsLowerStr(l int) string {
	str := Digits + AsciiLowercase
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomStr(l int) string {
	bytes := []byte(AsciiLetters)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomDigitsStr(l int) string {
	bytes := []byte(Digits + AsciiLowercase)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomUpperStr(l int) string {
	bytes := []byte(AsciiUppercase)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomLowerStr(l int) string {
	bytes := []byte(AsciiLowercase)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
