/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_channel_sdk
*/

package model

import "github.com/jinzhu/gorm"

type SysChannelSdk struct {
	gorm.Model
	ChannelName string `json:"channel_name" gorm:"not null;comment:'渠道名'"`
	SdkName     string `json:"sdk_name" gorm:"not null;comment:'SDK名称'"`
	SdkVersion  string `json:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}
