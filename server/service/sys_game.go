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

func ListGame() (err error, games []string) {
	err = global.GVA_DB.Model(&model.SysGame{}).Pluck("game_group", &games).Error
	return err, games
}

func GetGid(gameGroup string) (err error, game model.SysGame) {
	//u.Password = crypto.MD5([]byte(u.Password))
	err = global.GVA_DB.Where("game_group = ?", gameGroup).First(&game).Error
	return
}
