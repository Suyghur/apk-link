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

func CreateGame(bean model.SysGame) (err error) {
	//根据版本去判断是否存在
	isExist := !errors.Is(global.GvaDb.Where("game_group = ?", bean.GameGroup).First(&model.SysGame{}).Error, gorm.ErrRecordNotFound)
	//isExist := !global.GvaDb.Where("game_group = ?", bean.GameGroup).First(&game).RecordNotFound()
	//isExist为true表明读到了，不能新建
	if isExist {
		return errors.New("该游戏项目已存在")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}

func SearchGame(bean request.ReqGameListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysGame{})
	var games []model.SysGame
	if bean.Gid != "" {
		db = db.Where("gid = ?", bean.Gid)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&games).Error
	return err, games, total
}

func GetGid(gameGroup string) (err error, game model.SysGame) {
	err = global.GvaDb.Where("game_group = ?", gameGroup).First(&game).Error
	return
}
