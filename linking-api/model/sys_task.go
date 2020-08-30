/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_task
*/

package model

import "time"

type SysTask struct {
	TaskId      uint `json:"task_id" gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	IsWhiteBag  int        `json:"is_white_bag" gorm:"not null;default:'0';comment:'是否白包'"`
	IsPluginSdk int        `json:"has_plugin_sdk" gorm:"not null;default:'0';comment:'是否媒体包'"`
	GameGroup   string     `json:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid         string     `json:"gid" gorm:"not null;comment:'游戏组ID'"`
	OriginBag   string     `json:"origin_bag"`
	Ext         string     `json:"ext" gorm:"comment:'扩展字段'"`
}
