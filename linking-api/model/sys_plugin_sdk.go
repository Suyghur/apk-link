/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_plugin_sdk
*/

package model

import "github.com/jinzhu/gorm"

type SysPluginSdk struct {
	gorm.Model
	PluginName  string `json:"plugin_name" gorm:"not null;comment:'插件名'"`
	SdkName     string `json:"sdk_name" gorm:"not null;comment:'SDK名称'"`
	SdkVersion  string `json:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}
