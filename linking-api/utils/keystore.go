/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : keystore
*/

package utils

import (
	"io/ioutil"
	"linking-api/global"
	"os/exec"
	"strconv"
	"strings"
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

//生成签名文件Keystore
func CreateKeystore(GroupName string) (keystore Keystore, err error) {
	pyGroupName := GetLowerPinYinInitials(GroupName)
	keystore.KeystoreName = pyGroupName + "_" + "hl.keystore"
	keystore.KeystorePassword = pyGroupName + GetRandomLowerStr(4) + "_hl" + strconv.Itoa(time.Now().Year())
	keystore.KeystoreAlias = "alias." + pyGroupName + "_hl" + strconv.Itoa(time.Now().Year())
	keystore.KeystoreAliasPassword = keystore.KeystorePassword
	global.GVA_LOG.Debug("keytool -genkey -alias " + keystore.KeystoreAlias + " -keyalg RSA -validity 20000 -keystore " + keystore.KeystoreName + " -storepass " + keystore.KeystorePassword + " -keypass " + keystore.KeystoreAliasPassword + " -dname CN=hl, OU=hl, O=hl, L=hl, ST=hl, C=hl")
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
}

//获取Keystore中的指纹证书
func GetKeystoreFingerprints(keystore *Keystore) (fingerprints Fingerprints, err error) {
	cmd := exec.Command("keytool", "-list", "-v", "-keystore", keystore.KeystoreName, "-storepass", keystore.KeystorePassword)

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
		//截取md5
		md5Pos := strings.Index(result, "MD5:  ") + 6
		md5Pos2 := strings.Index(result[md5Pos:], "\n") - 6
		fingerprints.FingerprintsMD5 = result[md5Pos:(md5Pos + md5Pos2)]

		//截取sha1
		sha1Pos := strings.Index(result, "SHA1: ") + 6
		sha1Pos2 := strings.Index(result[sha1Pos:], "\n") - 6
		fingerprints.FingerprintsSHA1 = result[sha1Pos:(sha1Pos + sha1Pos2)]

		//截取sha256
		sha256Pos := strings.Index(result, "SHA256: ") + 8
		sha256Pos2 := strings.Index(result[sha256Pos:], "\n") - 8
		fingerprints.FingerprintsSHA256 = result[sha256Pos:(sha256Pos + sha256Pos2)]
	}

	return fingerprints, nil
}
