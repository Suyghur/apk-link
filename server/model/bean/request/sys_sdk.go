/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package request

type ReqSdkBean struct {
	SdkName     string `json:"sdk_name" form:"sdk_name" gorm:"not null;comment:'SDK名称'"`
	SdkVersion  string `json:"sdk_version" form:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" form:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" form:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}

type ReqFuseSdkBean struct {
	ReqSdkBean
}
type ReqListFuseSdkBean struct {
	ReqFuseSdkBean
	PageInfo
}

type ReqChannelSdkBean struct {
	ReqSdkBean
	ChannelName string `json:"channel_name" form:"channel_name" gorm:"not null;comment:'渠道名'"`
}

type ReqChannelSdksBean struct {
	ChannelName string `json:"channel_name" form:"channel_name" gorm:"not null;comment:'渠道名'"`
}

type ReqListChannelSdkBean struct {
	ReqChannelSdkBean
	PageInfo
}

type ReqPluginSdkBean struct {
	ReqSdkBean
	PluginName string `json:"plugin_name" form:"plugin_name" gorm:"not null;comment:'插件名'"`
}

type ReqPluginSdksBean struct {
	PluginName string `json:"plugin_name" form:"plugin_name" gorm:"not null;comment:'插件名'"`
}

type ReqListPluginSdkBean struct {
	ReqPluginSdkBean
	PageInfo
}
