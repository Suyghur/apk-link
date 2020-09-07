/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_plugin
*/

package model

import "github.com/jinzhu/gorm"

type SysPlugin struct {
	gorm.Model
	PluginName  string `json:"plugin_name" gorm:"not null;index:idx_plugin_name;comment:'插件名'"`
	PluginAlias string `json:"plugin_alias" gorm:"not null;index:idx_plugin_alias;comment:'插件别名'"`
}
