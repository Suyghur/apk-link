/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_fuse_sdk
*/

package model

import "gorm.io/gorm"

type SysFuseSdk struct {
	gorm.Model
	SdkVersion  string `json:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}
