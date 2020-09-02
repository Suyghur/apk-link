/*
@Time : 9/2/2020
@Author : #Suyghur,
@File : sys_script
*/

package service

import (
	"errors"
	"linking-api/global"
	"linking-api/model"
)

func CreateScript(s model.SysScript) (err error, scriptInter model.SysScript) {
	var script model.SysScript
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("game_group = ?", s.GameGroup).First(&script).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏已存在脚本文件"), scriptInter
	} else {
		err = global.GVA_DB.Create(s).Error
	}
	return err, s
}
