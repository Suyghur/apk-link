/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package service

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/bean/request"
)

func CreateGame(bean model.SysGame) (err error) {
	var game model.SysGame
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("game_group = ?", bean.GameGroup).First(&game).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏项目已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func SearchGame(bean request.ReqGameListBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysGame{})
	var games []model.SysGame
	if bean.Gid != "" {
		db = db.Where("gid = ?", bean.Gid)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&games).Error
	return err, games, total
}

func GetGid(gameGroup string) (err error, game model.SysGame) {
	err = global.GVA_DB.Where("game_group = ?", gameGroup).First(&game).Error
	return
}
