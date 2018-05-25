package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"microBlog/DB"
	"gopkg.in/mgo.v2/bson"
)

//登陆
func Login(c *gin.Context) {
	fmt.Println("this is a middleware!")
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Home(c *gin.Context) {
	tempArray:=[]int{1,2,3,4,5,6}
	articleList:=DB.Query("articles",bson.M{"title":"hello"})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articleList":articleList,
		"typeList":tempArray,
	})
}
