/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package model

import "github.com/jinzhu/gorm"

type SysSdk struct {
	gorm.Model
	SdkName     string `json:"sdk_name" gorm:"not null;comment:'SDK名称'"`
	SdkVersion  string `json:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
	SdkFileMD5  string `json:"sdk_file_md5" gorm:"not null;comment:'SDK文件MD5值'"`
}
