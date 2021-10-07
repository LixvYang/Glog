package routes

import (
	"glog/controller"
	"glog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//user api
		router.POST("user/add", controller.AddUser)
		router.GET("users", controller.GetUser)
		router.PUT("user/:id", controller.ExitUser)
		router.DELETE("user/:id", controller.DeleteUser)

		//category api
		router.GET("category", controller.GetCate)
		router.POST("category/add", controller.AddCategory)
		router.PUT("category/:id", controller.EditCate)
		router.DELETE("category/:id", controller.DeleteCate)

		//article api
		router.GET("article/info/:id", controller.GetArtInfo)
		router.GET("article", controller.GetArt)
		router.POST("article/add", controller.AddArticle)
		router.PUT("article/:id", controller.EditArt)
		router.DELETE("article/:id", controller.DeleteArt)
		router.GET("article/list/:id", controller.GetCateArt)
	}

	r.Run(utils.HttpPort)
}
