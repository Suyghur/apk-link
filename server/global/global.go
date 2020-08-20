/*
@Time : 2020/8/20
@Author : #Suyghur,
@File : global
*/
package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"server/config"
)

var (
	TP_DB     *gorm.DB
	TP_REDIS  *redis.Client
	TP_CONFIG config.Server
	TP_VP     *viper.Viper
	TP_LOG    *oplogging.Logger
)
