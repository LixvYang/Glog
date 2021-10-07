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
		//user api
		router.POST("user/add",controller.AddUser)
		router.GET("users",controller.GetUser)
		router.PUT("user/:id",controller.ExitUser)
		router.DELETE("user/:id",controller.DeleteUser)
	
		//category api
		router.GET("category",controller.GetCate)
		router.POST("category/add",controller.AddCategory)
		router.PUT("category/:id",controller.EditCate)
		router.DELETE("category/:id",controller.DeleteCate)
	}

	
	r.Run(utils.HttpPort)
}