/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_keystore
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
	"server/model/bean/response"
	"server/utils"
	"server/utils/aliyun"
	"os"
)

func SearchKeystore(bean request.ReqKeystoreListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysKeystore{})
	var keystores []model.SysKeystore
	if bean.GameGroup != "" {
		db = db.Where("game_group = ?", bean.GameGroup)
	}
	if bean.KeystoreName != "" {
		db = db.Where("keystore_name = ?", bean.KeystoreName)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&keystores).Error
	return err, keystores, total
}

func GetKeystores(gameGroup string) (err error, keystores []response.KeystoresResponse) {
	err = global.GvaDb.Model(&model.SysKeystore{}).Select("keystore_name , keystore_password , keystore_alias , keystore_alias_password , keystore_file_url").Where("game_group = ?", gameGroup).Scan(&keystores).Error
	return err, keystores

}

func CreateKeystore(gameGroup string) (err error) {
	var keystore model.SysKeystore
	pyGroupName := utils.GetLowerPinYinInitials(gameGroup)
	keystoreName := pyGroupName + "_" + "hl.keystore"
	//判断签名文件是否已经存在
	isExit := !errors.Is(global.GvaDb.Where("keystore_name = ?", keystoreName).First(&model.SysKeystore{}).Error, gorm.ErrRecordNotFound)
	//isExit := !global.GvaDb.Where("keystore_name = ?", keystoreName).First(&keystore).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏已存在签名文件")
	} else {
		newKeystore, err := utils.CreateKeystore(gameGroup)
		if err != nil {
			return err
		}
		fingerprints, err := utils.GetKeystoreFingerprints(&newKeystore)
		if err != nil {
			return err
		}
		//TODO 上传
		aliyun.Upload("linking/keystore/"+keystoreName, keystoreName)
		//TODO 删除
		err = os.Remove(keystoreName)
		if err != nil {
			return err
		}
		keystore.GameGroup = gameGroup
		keystore.KeystoreName = newKeystore.KeystoreName
		keystore.KeystorePassword = newKeystore.KeystorePassword
		keystore.KeystoreAlias = newKeystore.KeystoreAlias
		keystore.KeystoreAliasPassword = newKeystore.KeystoreAliasPassword
		keystore.FingerprintsMD5 = fingerprints.FingerprintsMD5
		keystore.FingerprintsSHA1 = fingerprints.FingerprintsSHA1
		keystore.FingerprintsSHA256 = fingerprints.FingerprintsSHA256
		keystore.KeystoreFileUrl = global.GvaConfig.Aliyun.Endpoint + "/linking/keystore/" + keystoreName
		err = global.GvaDb.Create(&keystore).Error

	}
	return err

}
