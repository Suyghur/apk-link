/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_task
*/

package model

import "time"

type SysTask struct {
	TaskId          uint `json:"task_id" gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
	IsFuseSdk       int        `json:"is_fuse_sdk" gorm:"default:'0';comment:'是否接入聚合SDK'"`
	IsWhiteBag      int        `json:"is_white_bag" gorm:"default:'0';comment:'是否白包'"`
	IsPluginSdk     int        `json:"has_plugin_sdk" gorm:"default:'0';comment:'是否媒体包'"`
	GameGroup       string     `json:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid             string     `json:"gid" gorm:"not null;comment:'游戏组ID'"`
	Aids            string     `json:"aids" gorm:"not null;comment:'子包AID'"`
	ChannelParams   string     `json:"channel_params" gorm:"comment:'渠道参数'"`
	PluginParams    string     `json:"plugin_params" gorm:"comment:'插件参数'"`
	GamePackageName string     `json:"game_package_name" gorm:"not null;comment:'游戏包名'"`
	GameName        string     `json:"game_name" gorm:"not null;comment:'游戏名'"`
	GameVersionCode int        `json:"game_version_code" gorm:"not null;comment:'游戏版本名'"`
	GameVersionName string     `json:"game_version_name" gorm:"not null;comment:'游戏版本名'"`
	GameOrientation int        `json:"game_orientation" gorm:"default:'0';comment:'游戏方向'"`
	GameIconUrl     string     `json:"game_icon_url" gorm:"游戏ICON链接"`
	GameLogoUrl     string     `json:"game_logo_url" gorm:"游戏LOGO链接"`
	GameSplashUrl   string     `json:"game_splash_url" gorm:"游戏闪屏链接"`
	GameLoginBgUrl  string     `json:"game_login_bg_url" gorm:"游戏登录背景链接"`
	Status          int        `json:"status" gorm:"default:'0';index:idx_status;gorm:'打包状态'"`
	//母包信息
	OriginBagName string `json:"origin_bag_name" gorm:"not null;comment:'母包名'"`
	OriginBagUrl  string `json:"origin_bag_url" gorm:"not null;comment:'母包链接'"`
	//聚合SDK信息
	FuseSdkFileName string `json:"fuse_sdk_file_name" gorm:"not null;comment:'聚合SDK名'"`
	FuseSdkVersion  string `json:"fuse_sdk_version" gorm:"not null;comment:'聚合SDK版本'"`
	FuseSdkUrl      string `json:"fuse_sdk_url" gorm:"not null;comment:'聚合SDK链接'"`
	//渠道SDK信息
	ChannelName        string `json:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	ChannelSdkFileName string `json:"channel_sdk_file_name" gorm:"not null;comment:'渠道SDK名'"`
	ChannelSdkVersion  string `json:"channel_sdk_version" gorm:"not null;comment:'渠道SDK版本'"`
	ChannelSdkUrl      string `json:"channel_sdk_url" gorm:"not null;comment:'渠道SDK链接'"`
	//插件SDK信息
	PluginName        string `json:"plugin_name" gorm:"index:idx_plugin_name;comment:'插件名'"`
	PluginSdkFileName string `json:"plugin_sdk_file_name" gorm:"comment:'插件SDK名'"`
	PluginSdkVersion  string `json:"plugin_sdk_version" gorm:"comment:'插件SDK版本'"`
	//游戏脚本信息
	GameScriptFileName string `json:"game_script_file_name" gorm:"comment:'游戏脚本名'"`
	GameScriptUrl      string `json:"game_script_url" gorm:"comment:'游戏脚本链接'"`
	//签名文件信息
	KeystoreName          string `json:"keystore_file_name" gorm:"not null;comment:'签名文件名'"`
	KeystorePassword      string `json:"keystore_password" gorm:"not null;comment:'签名文件密码'"`
	KeystoreAlias         string `json:"keystore_alias" gorm:"not null;comment:'签名文件别名'"`
	KeystoreAliasPassword string `json:"keystore_alias_password" gorm:"not null;comment:'签名文件别名密码'"`
	Ext                   string `json:"ext" gorm:"comment:'扩展字段'"`
}
