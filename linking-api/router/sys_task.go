/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_task
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "linking-api/api/v1"
)

func InitTaskRouter(r *gin.RouterGroup) {
	taskRouter := r.Group("task")
	{
		taskRouter.POST("list", v1.ListTask)
		taskRouter.POST("create", v1.CreateTask)
		taskRouter.POST("modify", v1.ModifyTask)
		taskRouter.POST("modify_status", v1.ModifyTaskStatus)
	}
}
