package controllers

import (
"fmt"
"github.com/gin-gonic/gin"
	"net/http"
)
//登陆
func Login(c *gin.Context) {
	fmt.Println("this is a middleware!")
	c.Redirect(http.StatusMovedPermanently,"/home")
}

func ToLogin(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}

//注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK,"register.html",nil)
}

func Home(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
}




