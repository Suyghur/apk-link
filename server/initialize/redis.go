/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : redis
*/

package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"server/global"
)

func Redis() {
	redisCfg := global.GvaConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GvaLog.Error("redis connect ping failed , err : ", zap.Any("err", err))
	} else {
		global.GvaLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.GvaRedis = client
	}
}
