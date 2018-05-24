package router

import (
	"github.com/gin-gonic/gin"
	"microBlog/controllers"
	"microBlog/config"
)

func init(){
	router:=gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "/editor")

	article:=router.Group("/article")
	{
		article.GET("/",controllers.Home)
		article.GET("/edit",controllers.Edit)
		article.GET("/publish",controllers.Publish)
		article.GET("/show",controllers.Show)
	}

	//用户操作
	router.GET("/home",controllers.Home)
	router.POST("/login",controllers.Login)
	router.GET("/toLogin",controllers.ToLogin)
	router.GET("/register",controllers.Register)
	router.Run(config.AppPort)
}
