/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package request

type ReqPluginBean struct {
	PluginName  string `json:"plugin_name" gorm:"not null;index:idx_plugin_name;comment:'插件名'"`
	PluginAlias string `json:"plugin_alias" gorm:"not null;index:idx_plugin_alias;comment:'插件别名'"`
}

type ReqPluginListBean struct {
	PluginAlias string `json:"plugin_alias" gorm:"not null;index:idx_plugin_alias;comment:'插件别名'"`
	PageInfo
}
