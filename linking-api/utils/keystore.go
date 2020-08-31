/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : keystore
*/

package utils

import (
	"io/ioutil"
	"linking-api/global"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Keystore struct {
	KeystoreName          string
	KeystorePassword      string
	KeystoreAlias         string
	KeystoreAliasPassword string
}
type Fingerprints struct {
	FingerprintsMD5    string
	FingerprintsSHA1   string
	FingerprintsSHA256 string
}

func CreateKeystore(GroupName string) (keystore Keystore, err error) {
	pyGroupName := getLowerPinYinInitials(GroupName)
	keystore.KeystoreName = pyGroupName + "_" + "hl.keystore"
	keystore.KeystorePassword = pyGroupName + GetRandomLowerStr(4) + "_hl" + strconv.Itoa(time.Now().Year())
	keystore.KeystoreAlias = "alias." + pyGroupName + "_hl" + strconv.Itoa(time.Now().Year())
	keystore.KeystoreAliasPassword = keystore.KeystorePassword
	cmd := exec.Command("keytool", "-genkey", "-alias", keystore.KeystoreAlias,
		"-keyalg", "RSA", "-validity", "20000", "-keystore", keystore.KeystoreName,
		"-storepass", keystore.KeystorePassword, "-keypass", keystore.KeystoreAliasPassword,
		"-dname", "CN=hl, OU=hl, O=hl, L=hl, ST=hl, C=hl")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		global.GVA_LOG.Error(err.Error())
		return keystore, err
	}
	if _, err = ioutil.ReadAll(stdout); err != nil {
		global.GVA_LOG.Error(err.Error())
		return keystore, err
	}
	return keystore, nil
	//global.GVA_LOG.Debug(err.Error())
	//pos := strings.Index(result, "MD5:  ")
	//pos2 := strings.Index(result[pos:], "\n")
	//fmt.Println(pos)
	//fmt.Println(pos2)
	//pos = pos + 6
	//pos2 = pos2 - 6
	//fmt.Println(result[(pos):(pos + pos2)])
}

func GetKeystoreFingerprints(keystore *Keystore) (fingerprints Fingerprints, err error) {
	//keytool -list -v -keystore yxz_hl.keystore
	workspace, _ := os.Getwd()
	cmd := exec.Command("keytool", "-list", "-v", "-keystore",
		workspace+"/assets/keystore/"+keystore.KeystoreName,
		"-storepass", keystore.KeystorePassword)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		global.GVA_LOG.Error(err.Error())
		return fingerprints, err
	}
	var optBytes []byte
	if optBytes, err = ioutil.ReadAll(stdout); err != nil {
		global.GVA_LOG.Error(err.Error())
		return fingerprints, err
	} else {
		result := string(optBytes)
		global.GVA_LOG.Debug(result)
	}

	return fingerprints, nil
}
