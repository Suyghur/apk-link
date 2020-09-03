/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : db_table
*/

package initialize

import (
	"linking-api/global"
	"linking-api/model"
)

func DBTables() {
	db := global.GVA_DB
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
	)
	global.GVA_LOG.Debug("register table success")
}
