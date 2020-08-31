/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_origin_bag
*/

package model

import "github.com/jinzhu/gorm"

type SysOriginBag struct {
	gorm.Model
	GameGroup       string `json:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid             string `json:"gid" gorm:"not null;comment:'游戏组ID'"`
	IsFuseSdk       int    `json:"is_fuse_sdk" gorm:"not null;default:'0';comment:'是否接入聚合SDK'"`
	SdkVersion      string `json:"sdk_version" gorm:"not null;index:idx_sdk_version;comment:'SDK版本'"`
	GameVersionCode int    `json:"game_version_code" gorm:"not null;comment:'游戏版本号'"`
	GameVersionName string `json:"game_version_name" gorm:"not null;comment:'游戏版本名'"`
	GameOrientation int    `json:"game_orientation" gorm:"default:'0';comment:'游戏方向'"`
	KeystoreName    string `json:"keystore_name" gorm:"not null;comment:'签名文件名'"`
	IconUrl         string `json:"icon_url" gorm:"comment:'ICON的URL'"`
	ApkUrl          string `json:"apk_url" gorm:"not null;comment:'APK的URL'"`
}
