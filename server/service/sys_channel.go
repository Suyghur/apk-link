/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
)

func CreateChannel(bean model.SysChannel) (err error) {
	//根据版本去判断是否存在
	isExist := !errors.Is(global.GvaDb.Where("cid = ?", bean.Cid, bean.ChannelName).First(&model.SysChannel{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		return errors.New("该渠道已存在")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}

func SearchChannel(bean request.ReqChannelListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysChannel{})
	var channels []model.SysChannel
	if bean.ChannelAlias != "" {
		db = db.Where("channel_alias = ?", bean.ChannelAlias)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&channels).Error
	return err, channels, total
}
