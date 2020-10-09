/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : cmd
*/

package utils

import (
	"go.uber.org/zap"
	"io/ioutil"
	"server/global"
	"os/exec"
)

func Command(cmdStr string) {
	cmd := exec.Command(cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		global.GvaLog.Error("stdout pipe err : ", zap.Any("err", err))
		return
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		global.GvaLog.Error("command err : ", zap.Any("err", err))
		return
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		global.GvaLog.Error("ioutil err : ", zap.Any("err", err))
		return
	} else {
		global.GvaLog.Info("opBytes", zap.Binary("op_bytes", opBytes))
	}

}
