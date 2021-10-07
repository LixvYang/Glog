package routes

import (
	"glog/controller"
	"glog/middleware"
	"glog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", controller.GetUsers)
		auth.PUT("user/:id", controller.EditUser)
		auth.DELETE("user/:id", controller.DeleteUser)
		// 分类模块的路由接口
		auth.GET("admin/category", controller.GetCate)
		auth.POST("category/add", controller.AddCategory)
		auth.PUT("category/:id", controller.EditCate)
		auth.DELETE("category/:id", controller.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", controller.GetArtInfo)
		auth.GET("admin/article", controller.GetArt)
		auth.POST("article/add", controller.AddArticle)
		auth.PUT("article/:id", controller.EditArt)
		auth.DELETE("article/:id", controller.DeleteArt)
	}

	router := r.Group("api/v1")
	{
		//user api
		router.POST("user/add", controller.AddUser)
		auth.GET("users", controller.GetUsers)
		//文章分类模块
		router.GET("category", controller.GetCate)
		router.GET("category/:id", controller.GetCateInfo)

		// 文章模块
		router.GET("article", controller.GetArt)
		router.GET("article/list/:id", controller.GetCateArt)
		router.GET("article/info/:id", controller.GetArtInfo)
		router.POST("login",controller.Login)

	}

	r.Run(utils.HttpPort)
}
