/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_options
*/

package service

import (
	"encoding/json"
	"io/ioutil"
	"server/global"
	"server/model"
	"server/model/bean/response"
	"net/http"
)

func GetOptions() (err error, options response.OptionsResponse) {
	tx := global.GVA_DB.Begin()
	err = tx.Model(&model.SysGame{}).Pluck("game_group", &options.GameOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	err = tx.Model(&model.SysChannel{}).Pluck("channel_name", &options.ChannelOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	err = tx.Model(&model.SysPlugin{}).Pluck("plugin_name", &options.PluginOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	tx.Commit()
	return err, options
}

func GetAids(gid string) (err error, aids []response.AidsResponse) {

	resp, err := http.Get("http://192.168.1.169:8000/getGameAidByGid.shtml?gid=" + gid)
	if err != nil || resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error(err.Error())
		return err, aids
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	type DataBean struct {
		Data []response.AidsResponse `json:"data"`
	}
	var bean DataBean
	if err = json.Unmarshal(s, &bean); err != nil {
		return err, aids
	}
	global.GVA_LOG.Debug(bean.Data)
	return err, bean.Data
}
