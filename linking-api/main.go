package main

import (
	"linking-api/core"
	"linking-api/global"
	"linking-api/initialize"
	"linking-api/utils"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()
	//结束前关闭数据块链接
	defer global.GVA_DB.Close()
	//if keystore, err := utils.CreateKeystore("游戏组"); err != nil {
	//	global.GVA_LOG.Error(err)
	//} else {
	//	global.GVA_LOG.Debug(keystore.KeystoreName)
	//	global.GVA_LOG.Debug(keystore.KeystorePassword)
	//	global.GVA_LOG.Debug(keystore.KeystoreAlias)
	//	global.GVA_LOG.Debug(keystore.KeystoreAliasPassword)
	//}
	keystore := utils.Keystore{
		KeystoreName:          "yxz_hl.keystore",
		KeystorePassword:      "yxzgnag_hl2020",
		KeystoreAlias:         "alias.yxz_hl2020",
		KeystoreAliasPassword: "yxzgnag_hl2020",
	}
	if _, err := utils.GetKeystoreFingerprints(&keystore); err != nil {
		global.GVA_LOG.Error(err)
	}

	core.RunServer()
}
