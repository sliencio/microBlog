package router

import (
	"github.com/gin-gonic/gin"
	"microBlog/controllers"
	"microBlog/config"
)

func init(){
	router:=gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/editor/css", "./static/editor/css")
	router.Static("/editor/js", "./static/editor/js")
	router.Static("/editor/lib", "./static/editor/lib")
	router.Static("/editor/fonts", "./static/editor/fonts")
	router.Static("/editor/images", "./static/editor/images")

	article:=router.Group("/article")
	{
		article.GET("/",controllers.Home)
		article.GET("/edit",controllers.Edit)
		article.POST("/publish",controllers.Publish)
		article.GET("/show",controllers.Show)
	}

	//用户操作
	router.GET("/home",controllers.Home)
	router.POST("/login",controllers.Login)
	router.GET("/toLogin",controllers.ToLogin)
	router.GET("/register",controllers.Register)
	router.Run(config.AppPort)
}
