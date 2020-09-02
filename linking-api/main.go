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
	//	if _, err := utils.GetKeystoreFingerprints(&keystore); err != nil {
	//		global.GVA_LOG.Error(err)
	//	}
	//}
	keystore := utils.Keystore{
		KeystoreName:          "yxz_hl.keystore",
		KeystorePassword:      "yxzsprr_hl2020",
		KeystoreAlias:         "alias.yxz_hl2020",
		KeystoreAliasPassword: "yxzsprr_hl2020",
	}
	if result, err := utils.GetKeystoreFingerprints(&keystore); err != nil {
		global.GVA_LOG.Error(err)
	} else {
		global.GVA_LOG.Debug(result.FingerprintsMD5)
		global.GVA_LOG.Debug(result.FingerprintsSHA1)
		global.GVA_LOG.Debug(result.FingerprintsSHA256)
	}
	//aliyun.Upload("linking/logo.png", "logo.png")
	core.RunServer()
}
