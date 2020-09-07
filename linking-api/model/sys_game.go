/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_game
*/

package model

import "github.com/jinzhu/gorm"

type SysGame struct {
	gorm.Model
	GameGroup string `json:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid       string `json:"gid" gorm:"not null;index:idx_gid;comment:'游戏组ID'"`
}
