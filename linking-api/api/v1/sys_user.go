/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : sys_user
*/

package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"linking-api/global"
	"linking-api/global/response"
	"linking-api/middleware"
	"linking-api/model"
	"linking-api/model/request"
	resp "linking-api/model/response"
	"linking-api/service"
	"linking-api/utils"
	"time"
)

func Test(c *gin.Context) {
	response.OkWithData(gin.H{"test": "test"}, c)
}

func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)

	UserVerify := utils.Rules{
		"username": {utils.NotEmpty()},
		"password": {utils.NotEmpty()},
		"nickname": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := &model.SysUser{
		Username: R.Username,
		Password: R.Password,
		Nickname: R.Nickname,
	}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.FailWithDetail(response.ERROR, fmt.Sprintf("%v", err), resp.SysUserResponse{User: userReturn}, c)
	} else {
		response.OkWithDetailed("注册成功", resp.SysUserResponse{User: userReturn}, c)
	}
}

func Login(c *gin.Context) {
	var L request.LoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &model.SysUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
	} else {
		tokenNext(c, *user)
	}
}

//登录成功后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{
		//唯一签名
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey),
	}
	clams := request.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.Nickname,
		Username: user.Username,
		//缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		BufferTime: 60 * 60 * 24,
		StandardClaims: jwt.StandardClaims{
			//签名生效时间
			NotBefore: time.Now().Unix() - 1000,
			//7天过期
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			//签名发行者
			Issuer: "Suyghur",
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
		return
	}
	err, jwtStr := service.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	} else if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(jwtStr, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}
