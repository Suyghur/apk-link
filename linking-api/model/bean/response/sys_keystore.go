/*
@Time : 2020/9/7
@Author : #Suyghur,
@File : sys_keystores
*/

package response

type KeystoresResponse struct {
	KeystoreName          string `json:"keystore_name" gorm:"not null;comment:'签名文件名'"`
	KeystorePassword      string `json:"keystore_password" gorm:"not null;comment:'签名文件密码'"`
	KeystoreAlias         string `json:"keystore_alias" gorm:"not null;comment:'签名文件别名'"`
	KeystoreAliasPassword string `json:"keystore_alias_password" gorm:"not null;comment:'签名文件别名密码'"`
	KeystoreFileUrl       string `json:"keystore_file_url" gorm:"not null;comment:'签名文件链接'"`
}
