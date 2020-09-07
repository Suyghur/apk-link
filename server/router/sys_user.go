/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : sys_user
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitUserRouter(r *gin.RouterGroup) {
	//UserRouter := Router.Group("user").
	//	Use(middleware.JWTAuth()).
	//	Use(middleware.CasbinHandler()).
	//	Use(middleware.OperationRecord())
	//{
	//	UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
	//	UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)   // 上传头像
	//	UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
	//	UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
	//	UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
	//}
	userRouter := r.Group("user").Use(middleware.JWTAuth())
	{
		userRouter.POST("changePassword", v1.ChangePassword)
		userRouter.POST("getUserInfo", v1.GetUserInfo)
	}
}
