/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : jwt_black_list
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"time"
)

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GvaDb.Create(&jwtList).Error
	return
}

func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GvaDb.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GvaRedis.Get(userName).Result()
	return err, redisJWT
}

func SetRedisJWT(jwt string, userName string) (err error) {
	//redis的过期时间等于jwt的过期时间
	timer := 60 * 60 * 24 * 7 * time.Second
	err = global.GvaRedis.Set(userName, jwt, timer).Err()
	return err
}
