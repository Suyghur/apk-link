/*
@Time : 2020/9/8
@Author : #Suyghur,
@File : sys_link
*/

package model

import "github.com/jinzhu/gorm"

type SysLink struct {
	gorm.Model
	TaskId          uint   `json:"task_id" gorm:"not null;index:idx_task_id;comment:'任务ID'"`
	GameGroup       string `json:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	GameName        string `json:"game_name" gorm:"not null;comment:'游戏名'"`
	GameVersionCode int    `json:"game_version_code" gorm:"not null;comment:'游戏版本号'"`
	GameVersionName string `json:"game_version_name" gorm:"not null;comment:'游戏版本名'"`
	Aid             string `json:"aid" gorm:"not null;comment:'子包AID'"`
	FuseSdkVersion  string `json:"fuse_sdk_version" gorm:"not null;comment:'聚合SDK版本'"`
	ChannelName     string `json:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	PluginName      string `json:"plugin_name" gorm:"index:idx_plugin_name;comment:'插件名'"`
	LinkUrl         string `json:"link_url" gorm:"comment:'分发链接'"`
}
