package routes

import (
	"glog/utils"
	"glog/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.POST("user/add",controller.AddUser)
		router.GET("users",controller.GetUser)
		router.PUT("user/:id",controller.ExitUser)
		router.DELETE("user/:id",controller.DeleteUser)
	
	}

	
	r.Run(utils.HttpPort)
}