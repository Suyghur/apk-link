/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

type ReqPluginBean struct {
	FormId      string `json:"form_id" form:"form_id" gorm:"not null;index:idx_form_id;comment:'分发ID'"`
	PluginName  string `json:"plugin_name" form:"plugin_name" gorm:"not null;index:idx_plugin_name;comment:'插件名'"`
	PluginAlias string `json:"plugin_alias" form:"plugin_alias" gorm:"not null;index:idx_plugin_alias;comment:'插件别名'"`
}

type ReqPluginListBean struct {
	PluginAlias string `json:"plugin_alias" form:"plugin_alias" gorm:"not null;index:idx_plugin_alias;comment:'插件别名'"`
	PageInfo
}
