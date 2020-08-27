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
		model.SysUser{},
	)
}
