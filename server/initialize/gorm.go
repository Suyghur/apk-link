/*
@Time : 10/9/20
@Author : #Suyghur,
@File : gorm
*/

package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"server/global"
	"server/model"
)

var err error

func Gorm() {
	switch global.GvaConfig.System.DbType {
	case "mysql":
		GormMysql()
	case "postgresql":
		GormPostgreSql()
	}
}

func GormDBTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysKeystore{},
		model.JwtBlacklist{},
		model.SysUser{},
		model.SysFuseSdk{},
		model.SysChannelSdk{},
		model.SysPluginSdk{},
		model.SysTask{},
		model.SysOrigin{},
		model.SysScript{},
		model.SysGame{},
		model.SysPlugin{},
		model.SysChannel{},
		model.SysLink{},
		model.SysGid{},
	)
	if err != nil {
		global.GvaLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GvaLog.Info("register table success")
}

func GormMysql() {
	m := global.GvaConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   //DSN data source name
		DefaultStringSize:         191,   //string 类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用 datetime 精度，Mysql 5.6 之前不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，Mysql 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //用 `change` 重命名列，Mysql 8 之前的数据块和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, //根据版本自动配置
	}
	gormConfig := config(m.LogMode)
	if global.GvaDb, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GvaLog.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		GormDBTables(global.GvaDb)
		sqlDB, _ := global.GvaDb.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	}
}

func GormPostgreSql() {
}

func config(mode bool) (c *gorm.Config) {
	if mode {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
