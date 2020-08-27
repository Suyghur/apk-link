/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : jwt_black_list
*/

package service

import (
	"linking-api/global"
	"linking-api/model"
	"time"
)

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

func IsBlacklist(jwt string, jwtList model.JwtBlacklist) bool {
	isNotFound := global.GVA_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

func SetRedisJWT(jwt string, userName string) (err error) {
	//redis的过期时间等于jwt的过期时间
	timer := 60 * 60 * 24 * 7 * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
