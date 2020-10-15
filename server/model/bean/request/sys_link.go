/*
@Time : 2020/9/8
@Author : #Suyghur,
@File : sys_link
*/

package request

type ReqLinkBean struct {
	TaskId         uint   `json:"task_id" gorm:"primary_key"`
	GameSite       string `json:"game_site" gorm:"not null;index:idx_game_site;comment:'游戏组'"`
	Gid            string `json:"gid" gorm:"not null;index:idx_gid;comment:'游戏ID'"`
	Cid            string `json:"cid" gorm:"not null;idx_cid;comment:'渠道ID'"`
	FormId         string `json:"form_id" gorm:"not null;index:idx_form_id;comment:'分发ID'"`
	Aid            string `json:"aid" gorm:"not null;comment:'包ID'"`
	GameName       string `json:"game_name" gorm:"not null;comment:'游戏名'"`
	FuseSdkVersion string `json:"fuse_sdk_version" gorm:"not null;comment:'聚合SDK版本'"`
	ChannelName    string `json:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	PluginName     string `json:"plugin_name" gorm:"index:idx_plugin_name;comment:'插件名'"`
	LinkUrl        string `json:"link_url" gorm:"comment:'分发链接'"`
}

type ReqLinkListBean struct {
	TaskId      uint   `json:"task_id" form:"task_id" gorm:"primary_key"`
	GameSite    string `json:"game_site" form:"game_site" gorm:"not null;index:idx_game_site;comment:'游戏组标识'"`
	Aid         string `json:"aid" form:"aid" gorm:"not null;comment:'子包AID'"`
	ChannelName string `json:"channel_name" form:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	PluginName  string `json:"plugin_name" form:"plugin_name" gorm:"index:idx_plugin_name;comment:'插件名'"`
	PageInfo
}
