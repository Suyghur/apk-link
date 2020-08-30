/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_plugin_sdk
*/

package model

type SysPluginSdk struct {
	SysSdk
	PluginName string `json:"plugin_name" gorm:"not null;comment:'插件名'"`
	//PluginDisplayName string `json:"plugin_display_name" gorm:"not null;comment:'插件展示名'"`
}
