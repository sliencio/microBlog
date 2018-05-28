package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"microBlog/DB"
	"microBlog/model"
	"net/http"
	"time"
)

//编辑
func Edit(c *gin.Context) {
	CheckLogin(c)
	temp:="选项1,选项2,选项3,选项4,选项5,选项6"
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"typeList":temp,
	})
	fmt.Println("this is a middleware!")
}

//展示文章列表
func Show(c *gin.Context) {
	CheckLogin(c)
	objID := c.Param("_id")
	fmt.Println("------:Id:" + objID)
	c.HTML(http.StatusOK, "showArticle.html", gin.H{"content": "dsfsdfsdf"})
}

//发布
func Publish(c *gin.Context) {
	article := &model.Article{
		Id:       bson.NewObjectId(),
		Uid:      bson.NewObjectId(),
		Title:    "hello",
		Time:     time.Now(),
		Classify: "go",
		Content:  "go is fine",
	}
	err := DB.Insert("articles", article)
	fmt.Println("err:", err)
	if err == nil {
		// c.Redirect(http.StatusMovedPermanently,"/home")
		c.JSON(http.StatusOK, gin.H{
			"ret": true,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"ret": false,
		})
	}
}
