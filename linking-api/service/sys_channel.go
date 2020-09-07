/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package service

import (
	"errors"
	"linking-api/global"
	"linking-api/model"
)

func CreateChannel(bean model.SysChannel) (err error) {
	var channel model.SysChannel
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("channel_name = ?", bean.ChannelName).First(&channel).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该渠道已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ListChannel() (err error, games []string) {
	err = global.GVA_DB.Model(&model.SysGame{}).Pluck("game_group", &games).Error
	return err, games
}
