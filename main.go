package main

import (
	_ "microBlog/router"
	_"microBlog/DataManager"
	"github.com/gin-gonic/gin"
)


func main() {
	//全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	gin.SetMode(gin.DebugMode)
}
