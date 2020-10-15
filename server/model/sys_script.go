/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_script
*/

package model

import "gorm.io/gorm"

type SysScript struct {
	gorm.Model
	Gid            string `json:"gid" gorm:"not null;comment:'游戏组ID'"`
	GameSite       string `json:"game_site" gorm:"not null;comment:'游戏标识'"`
	ScriptFileName string `json:"script_file_name" gorm:"not null;comment:'脚本文件名'"`
	ScriptFileUrl  string `json:"script_file_url" gorm:"not null;comment:'脚本文件链接'"`
}
