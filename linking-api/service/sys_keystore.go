/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_keystore
*/

package service

import (
	"linking-api/global"
	"linking-api/model"
	"linking-api/model/request"
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

//func CreateKeystore(gameGroup string) (err error) {
//
//}
