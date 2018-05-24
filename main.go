package main

import (
	"github.com/gin-gonic/gin"
	_"microBlog/router"
)

func main() {
	//全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	gin.SetMode(gin.DebugMode)
}
