/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_task
*/

package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

func InitTaskRouter(r *gin.RouterGroup) {
	taskRouter := r.Group("task").Use(middleware.JWTAuth())
	{
		taskRouter.POST("searchTasks", v1.SearchTasks)
		taskRouter.POST("getTaskInfo", v1.GetTaskInfo)
		taskRouter.POST("createTask", v1.CreateTask)
		taskRouter.POST("modifyTask", v1.ModifyTask)
		taskRouter.POST("deleteTask", v1.DeleteTask)
		taskRouter.POST("modifyStatus", v1.ModifyTaskStatus)
	}
}
