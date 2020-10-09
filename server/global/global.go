/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : global
*/
package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	GvaDb     *gorm.DB
	GvaRedis  *redis.Client
	GvaConfig config.Server
	GvaVp     *viper.Viper
	GvaLog    *zap.Logger
	//GvaLog    *oplogging.Logger
)
