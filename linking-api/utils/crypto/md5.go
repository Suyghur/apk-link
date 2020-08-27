/*
@Time : 8/28/2020
@Author : #Suyghur,
@File : md5
*/

package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
