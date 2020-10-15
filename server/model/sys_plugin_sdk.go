/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_plugin_sdk
*/

package model

import "gorm.io/gorm"

type SysPluginSdk struct {
	gorm.Model
	FormId      string `json:"form_id" form:"form_id" gorm:"not null;index:idx_form_id;comment:'分发ID'"`
	PluginName  string `json:"plugin_name" form:"plugin_name" gorm:"not null;index:idx_plugin_name;comment:'插件名'"`
	PluginAlias string `json:"plugin_alias" form:"plugin_alias" gorm:"not null;comment:'SDK名称'"`
	SdkVersion  string `json:"sdk_version" form:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" form:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" form:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}
