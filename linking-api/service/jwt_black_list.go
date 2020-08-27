/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : jwt_black_list
*/

package service

import (
	"linking-api/global"
	"linking-api/model"
)

func IsBlacklist(jwt string, jwtList model.JwtBlacklist) bool {
	isNotFound := global.GVA_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}
