/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : db_table
*/

package initialize

import (
	"server/global"
	"server/model"
)

func DBTables() {
	db := global.GvaDb
	db.AutoMigrate(
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
	)
	global.GvaLog.Debug("register table success")
}
