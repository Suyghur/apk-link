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

func CreateGid(bean model.SysGame) (err error) {
	//根据版本去判断是否存在
	isExist := !errors.Is(global.GvaDb.Where("gid = ?", bean.Gid).First(&model.SysGame{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		return errors.New("该游戏ID已存在")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}

func CreateGameSite(bean model.SysGame) (err error) {
	//根据版本去判断是否存在
	isExist := !errors.Is(global.GvaDb.Where("game_site = ?", bean.GameSite).First(&model.SysGame{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		return errors.New("该游戏项目已存在")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}

func SearchGameSite(bean request.ReqGameListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysGame{})
	var games []model.SysGame
	if bean.GameSite != "" {
		db = db.Where("gid = ?", bean.GameSite)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&games).Error
	return err, games, total
}

func GetGids(gameSite string) (err error, gids []string) {
	err = global.GvaDb.Model(&model.SysGid{}).Select("gid").Where("game_site = ?", gameSite).Scan(&gids).Error
	return err, gids
}
