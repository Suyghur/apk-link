/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : cmd
*/

package utils

import (
	"io/ioutil"
	"server/global"
	"os/exec"
)

func Command(cmdStr string) {
	cmd := exec.Command(cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		global.GVA_LOG.Error(err)
		return
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		global.GVA_LOG.Error(err)
		return
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		global.GVA_LOG.Error(err)
		return
	} else {
		global.GVA_LOG.Info(opBytes)
	}

}
