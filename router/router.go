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
		article.GET("/show/:articleId",controllers.Show)
		article.GET("/delete/:articleId",controllers.Delete)
		article.GET("/reEdit/:articleId",controllers.ReEdit)
		article.GET("/example",controllers.Example)
	}

	//用户操作
	user:=router.Group("/user")
	{
		user.POST("/login",controllers.Login)
		user.GET("/toLogin",controllers.ToLogin)
		user.POST("/register",controllers.Register)
		user.GET("/toRegister",controllers.ToRegister)
	}
	router.GET("/",controllers.Home)
	router.GET("/home",controllers.Home)

	router.Run(config.AppPort)
}
