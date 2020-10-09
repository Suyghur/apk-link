/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package request

type ReqOriginBean struct {
	GameGroup       string `json:"game_group" form:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	Gid             string `json:"gid" form:"gid" gorm:"not null;comment:'游戏组ID'"`
	IsFuseSdk       int    `json:"is_fuse_sdk" form:"is_fuse_sdk" gorm:"default:'1';comment:'是否接入聚合SDK'"`
	SdkVersion      string `json:"sdk_version" form:"sdk_version" gorm:"not null;index:idx_sdk_version;comment:'聚合SDK版本'"`
	GameFileName    string `json:"game_file_name" form:"game_file_name" gorm:"not null;index:idx_game_file_name;comment:'母包文件名'"`
	GameVersionCode int    `json:"game_version_code" form:"game_version_code" gorm:"not null;comment:'游戏版本号'"`
	GameVersionName string `json:"game_version_name" form:"game_version_name" gorm:"not null;comment:'游戏版本名'"`
	GameOrientation int    `json:"game_orientation" form:"game_orientation" gorm:"default:'0';comment:'游戏方向'"`
	KeystoreName    string `json:"keystore_name" form:"keystore_name" gorm:"not null;comment:'签名文件名'"`
	IconUrl         string `json:"icon_url" form:"icon_url" gorm:"comment:'ICON的URL'"`
	ApkUrl          string `json:"apk_url" form:"apk_url" gorm:"not null;comment:'APK的URL'"`
}

type ReqSearchOriginBean struct {
	GameGroup  string `json:"game_group" form:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
	SdkVersion string `json:"sdk_version" form:"sdk_version" gorm:"not null;index:idx_sdk_version;comment:'聚合SDK版本'"`
	PageInfo
}

type ReqGetOriginListBean struct {
	GameGroup string `json:"game_group" form:"game_group" gorm:"not null;index:idx_game_group;comment:'游戏组'"`
}
