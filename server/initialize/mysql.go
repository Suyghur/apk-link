/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : mysql
*/

package initialize

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	//admin := global.GvaConfig.Mysql
	//if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
	//	global.GvaLog.Error("MySQL启动异常", err)
	//	os.Exit(0)
	//} else {
	//	global.GvaDb = db
	//	global.GvaDb.DB().SetMaxIdleConns(admin.MaxIdleConns)
	//	global.GvaDb.DB().SetMaxOpenConns(admin.MaxOpenConns)
	//	global.GvaDb.DB().SetConnMaxLifetime(time.Second * 30)
	//
	//	global.GvaDb.LogMode(admin.LogMode)
	//}
}
