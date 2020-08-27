/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : redis
*/

package initialize

import (
	"github.com/go-redis/redis"
	"linking-api/global"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error(err)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
