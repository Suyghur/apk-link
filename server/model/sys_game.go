/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_game
*/

package model

import "gorm.io/gorm"

type SysGame struct {
	gorm.Model
	Gid      string `json:"gid" form:"gid" gorm:"not null;index:idx_gid;comment:'游戏ID'"`
	GameSite string `json:"game_site" gorm:"not null;index:idx_game_site;comment:'游戏组标识'"`
}
