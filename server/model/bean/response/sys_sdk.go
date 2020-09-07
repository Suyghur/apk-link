/*
@Time : 2020/9/7
@Author : #Suyghur,
@File : sys_sdk
*/

package response

type SdksResponse struct {
	SdkVersion  string `json:"sdk_version" gorm:"not null;comment:'SDK版本'"`
	SdkFileName string `json:"sdk_file_name" gorm:"not null;comment:'SDK文件名'"`
	SdkFileUrl  string `json:"sdk_file_url" gorm:"not null;comment:'SDK文件链接'"`
}

