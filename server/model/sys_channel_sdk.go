/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_channel_sdk
*/

package model

import "gorm.io/gorm"

type SysChannelSdk struct {
	gorm.Model
	Cid          string `json:"cid" form:"cid" gorm:"not null;index:idx_cid;comment:'渠道ID'"`
	ChannelName  string `json:"channel_name" form:"channel_name" gorm:"not null;comment:'渠道名'"`
	ChannelAlias string `json:"channel_alias" form:"channel_alias" gorm:"not null;comment:'渠道标识'"`
	SdkVersion   string `json:"sdk_version" form:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName  string `json:"sdk_file_name" form:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl   string `json:"sdk_file_url" form:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}
