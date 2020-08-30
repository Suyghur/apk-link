/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_keystore
*/

package model

import "github.com/jinzhu/gorm"

type SysKeystore struct {
	gorm.Model
	KeystoreName          string `json:"keystore_name" gorm:"comment:'签名文件名'"`
	KeystorePassword      string `json:"keystore_password" gorm:"comment:'签名文件密码'"`
	KeystoreAlias         string `json:"keystore_alias" gorm:"comment:'签名文件别名'"`
	KeystoreAliasPassword string `json:"keystore_alias_password" gorm:"comment:'签名文件别名密码'"`
	FingerprintsMD5       string `json:"fingerprints_md5" gorm:"comment:'签名文件指纹MD5值'"`
	FingerprintsSHA1      string `json:"fingerprints_sha1" gorm:"comment:'签名文件指纹SHA1值'"`
	FingerprintsSHA256    string `json:"fingerprints_sha256" gorm:"comment:'签名文件指纹SHA256值'"`
	KeystoreFileUrl       string `json:"keystore_file_url" gorm:"comment:'签名文件链接'"`
	KeystoreFileMD5       string `json:"keystore_file_md5" gorm:"comment:'签名文件MD5值'"`
}
