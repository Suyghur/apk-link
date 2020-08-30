/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_channel_sdk
*/

package model

type SysChannelSdk struct {
	SysSdk
	ChannelName string `json:"channel_name" gorm:"comment:'渠道名'"`
}
