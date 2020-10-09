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
	"server/global"
	"server/global/response"
	"server/middleware"
	"server/model"
	"server/model/bean/request"
	resp "server/model/bean/response"
	"server/service"
	"server/utils"
	"time"
)

func Register(c *gin.Context) {
	var reqStruct request.RegisterBean
	_ = c.ShouldBindJSON(&reqStruct)

	userVerify := utils.Rules{
		"username": {utils.NotEmpty()},
		"password": {utils.NotEmpty()},
		"nickname": {utils.NotEmpty()},
	}
	userVerifyErr := utils.Verify(reqStruct, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
		return
	}
	user := &model.SysUser{
		Username: reqStruct.Username,
		Password: reqStruct.Password,
		Nickname: reqStruct.Nickname,
	}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.FailWithDetail(response.ERROR, fmt.Sprintf("%v", err), resp.SysUserResponse{User: userReturn}, c)
	} else {
		response.OkWithDetailed("注册成功", resp.SysUserResponse{User: userReturn}, c)
	}
}

func Login(c *gin.Context) {
	var reqStruct request.LoginBean
	_ = c.ShouldBindJSON(&reqStruct)
	userVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	userVerifyErr := utils.Verify(reqStruct, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
		return
	}
	U := &model.SysUser{Username: reqStruct.Username, Password: reqStruct.Password}
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
		SigningKey: []byte(global.GvaConfig.JWT.SigningKey),
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
	if !global.GvaConfig.System.UseMultipoint {
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

func GetUserInfo(c *gin.Context) {
	var bean request.UserInfoBean
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"Username": {utils.NotEmpty()},
	}
	verifyErr := utils.Verify(bean, verifyRules)
	if verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, userReturn := service.Info(bean.Username)
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
	} else {
		response.OkWithDetailed("获取用户信息成功", gin.H{
			"username": userReturn.Username,
			"uuid":     userReturn.UUID,
			"nickname": userReturn.Nickname,
		}, c)
	}
}

func ChangePassword(c *gin.Context) {
	var bean request.ChangePasswordBean
	_ = c.ShouldBindJSON(&bean)
	userVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"NewPassword": {utils.NotEmpty()},
	}
	verifyErr := utils.Verify(bean, userVerify)
	if verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	user := &model.SysUser{
		Username: bean.Username,
		Password: bean.Password,
	}

	if err, _ := service.ChangePassword(user, bean.NewPassword); err != nil {
		response.FailWithMessage("修改失败，请检查用户名密码", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
