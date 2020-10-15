/*
@Time : 2020/9/8
@Author : #Suyghur,
@File : sys_link
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
	"strings"
)

func SearchLink(bean request.ReqLinkListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysLink{})
	var links []model.SysLink
	if bean.TaskId > 0 {
		db = db.Where("task_id = ?", bean.TaskId)
	}
	if bean.GameSite != "" {
		db = db.Where("game_site = ?", bean.GameSite)
	}
	if bean.Aid != "" {
		db = db.Where("aid = ?", bean.Aid)
	}
	if bean.ChannelName != "" {
		db = db.Where("channel_name = ?", bean.ChannelName)
	}
	if bean.PluginName != "" {
		db = db.Where("plugin_name = ?", bean.PluginName)
	}
	db = db.Where("link_url <> ''")
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&links).Error
	return err, links, total
}

func CreateLink(taskId uint) (err error) {
	//判断游戏脚本是否存在
	isExit := !errors.Is(global.GvaDb.Where("task_id = ?", taskId).First(&model.SysLink{}).Error, gorm.ErrRecordNotFound)
	//isExit := !global.GvaDb.Where("task_id = ?", taskId).First(&link).RecordNotFound()
	if isExit {
		return errors.New("该任务已存在")
	} else {
		type taskBean struct {
			GameSite        string
			GameName        string
			GameVersionCode int
			GameVersionName string
			Aids            string
			FuseSdkVersion  string
			ChannelName     string
			PluginName      string
		}
		var bean taskBean
		if err = global.GvaDb.Model(&model.SysTask{}).Select("game_site , game_name , game_version_code , game_version_name , aids , fuse_sdk_version , channel_name , plugin_name").Where("task_id = ?", taskId).Scan(&bean).Error; err != nil {
			return err
		} else {
			tx := global.GvaDb.Begin()
			aidsArr := strings.Split(bean.Aids, ",")
			for _, value := range aidsArr {
				link := model.SysLink{
					TaskId:          taskId,
					GameSite:        bean.GameSite,
					GameName:        bean.GameName,
					GameVersionCode: bean.GameVersionCode,
					GameVersionName: bean.GameVersionName,
					Aid:             value,
					FuseSdkVersion:  bean.FuseSdkVersion,
					ChannelName:     bean.ChannelName,
					PluginName:      bean.PluginName,
					LinkUrl:         "",
				}
				if err = tx.Create(&link).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
			tx.Commit()
		}
	}
	return err
}

func ReportLink(bean model.SysLink) (err error) {
	var link model.SysLink
	//判断游戏脚本是否存在
	isExist := !errors.Is(global.GvaDb.Where("task_id = ?", bean.TaskId).First(&model.SysLink{}).Error, gorm.ErrRecordNotFound)
	//isExist := !global.GvaDb.Where("task_id = ?", bean.TaskId).First(&link).RecordNotFound()
	//isExist为true表明读到了，不能新建
	if isExist {
		err = global.GvaDb.Where("task_id = ?", bean.TaskId).First(&link).Updates(map[string]interface{}{"link_url": bean.LinkUrl}).Error
	} else {
		return errors.New("没有可上报的任务")
	}
	return err
}
