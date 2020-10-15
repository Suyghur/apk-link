/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_origin_bag
*/

package model

import "gorm.io/gorm"

type SysOrigin struct {
	gorm.Model
	Gid             string `json:"gid" form:"gid" gorm:"not null;index:idx_gid;comment:'游戏组ID'"`
	GameSite        string `json:"game_site" form:"game_site" gorm:"not null;index:idx_game_site;comment:'游戏标识'"`
	SdkVersion      string `json:"sdk_version" form:"sdk_version" gorm:"not null;index:idx_sdk_version;comment:'SDK版本'"`
	GameFileName    string `json:"game_file_name" form:"game_file_name" gorm:"not null;index:idx_game_file_name;comment:'母包文件名'"`
	GameVersionCode int    `json:"game_version_code" form:"game_version_code" gorm:"not null;comment:'游戏版本号'"`
	GameVersionName string `json:"game_version_name" form:"game_version_name" gorm:"not null;comment:'游戏版本名'"`
	GameOrientation int    `json:"game_orientation" form:"game_orientation" gorm:"default:0;comment:'游戏方向'"`
	KeystoreName    string `json:"keystore_name" form:"keystore_name" gorm:"not null;comment:'签名文件名'"`
	ApkUrl          string `json:"apk_url" form:"apk_url" gorm:"not null;comment:'APK的URL'"`
}
