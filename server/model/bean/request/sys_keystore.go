/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_keystore
*/

package request

type ReqKeystoreBean struct {
	GameGroup             string `json:"game_group" form:"game_group" gorm:"not null;comment:'游戏组'"`
	KeystoreName          string `json:"keystore_name" form:"keystore_name" gorm:"not null;comment:'签名文件名'"`
	KeystorePassword      string `json:"keystore_password" form:"keystore_password" gorm:"not null;comment:'签名文件密码'"`
	KeystoreAlias         string `json:"keystore_alias" form:"keystore_alias" gorm:"not null;comment:'签名文件别名'"`
	KeystoreAliasPassword string `json:"keystore_alias_password" form:"keystore_alias_password" gorm:"not null;comment:'签名文件别名密码'"`
	FingerprintsMD5       string `json:"fingerprints_md5" form:"fingerprints_md5" gorm:"not null;comment:'签名文件指纹MD5值'"`
	FingerprintsSHA1      string `json:"fingerprints_sha1" form:"fingerprints_sha1" gorm:"not null;comment:'签名文件指纹SHA1值'"`
	FingerprintsSHA256    string `json:"fingerprints_sha256" form:"fingerprints_sha256" gorm:"not null;comment:'签名文件指纹SHA256值'"`
	KeystoreFileUrl       string `json:"keystore_file_url" form:"keystore_file_url" gorm:"not null;comment:'签名文件链接'"`
}

type ReqKeystoreListBean struct {
	ReqKeystoreBean
	PageInfo
}
