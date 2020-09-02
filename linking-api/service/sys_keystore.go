/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_keystore
*/

package service

import (
	"errors"
	"linking-api/global"
	"linking-api/model"
	"linking-api/model/request"
	"linking-api/utils"
)

func ListKeystore(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysKeystore{})
	var keystoreList []model.SysKeystore
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("keystore_name").Find(&keystoreList).Error
	return err, keystoreList, total
}

func CreateKeystore(gameGroup string) (err error, keystoreInter model.SysKeystore) {
	var keystore model.SysKeystore
	pyGroupName := utils.GetLowerPinYinInitials(gameGroup)
	keystoreName := pyGroupName + "_" + "hl.keystore"
	//判断签名文件是否已经存在
	isExit := !global.GVA_DB.Where("keystore_name = ?", keystoreName).First(&keystore).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏已存在签名文件"), keystoreInter
	} else {
		newKeystore, err := utils.CreateKeystore(gameGroup)
		if err != nil {
			return err, keystoreInter
		}
		fingerprints, err := utils.GetKeystoreFingerprints(&newKeystore)
		if err != nil {
			return err, keystoreInter
		}
		keystore.KeystoreName = newKeystore.KeystoreName
		keystore.KeystorePassword = newKeystore.KeystorePassword
		keystore.KeystoreAlias = newKeystore.KeystoreAlias
		keystore.KeystoreAliasPassword = newKeystore.KeystoreAliasPassword
		keystore.FingerprintsMD5 = fingerprints.FingerprintsMD5
		keystore.FingerprintsSHA1 = fingerprints.FingerprintsSHA1
		keystore.FingerprintsSHA256 = fingerprints.FingerprintsSHA256
		err = global.GVA_DB.Create(&keystore).Error
	}
	return err, keystore

}
