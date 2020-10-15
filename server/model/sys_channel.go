/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_channel
*/

package model

import "gorm.io/gorm"

type SysChannel struct {
	gorm.Model
	Cid          string `json:"cid" gorm:"not null;index:idx_cid;comment:'渠道ID'"`
	ChannelName  string `json:"channel_name" gorm:"not null;index:idx_channel_name;comment:'渠道名'"`
	ChannelAlias string `json:"channel_alias" gorm:"not null;index:idx_channel_alias;comment:'渠道别名'"`
}
