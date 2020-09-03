/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : jwt
*/

package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	UUID     uuid.UUID
	ID       uint
	NickName string
	Username string
	//AuthorityId string
	BufferTime int64
	jwt.StandardClaims
}
