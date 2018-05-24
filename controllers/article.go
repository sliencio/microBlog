package controllers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"microBlog/DB"
	"github.com/gin-gonic/gin"
	"net/http"
)

//编辑
func Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.html", nil)
	fmt.Println("this is a middleware!")
}

//展示文章列表
func Show(c *gin.Context) {
	objID := c.Param("_id")
	fmt.Println("------:Id:" + objID)
	c.HTML(http.StatusOK, "showArticle.html", gin.H{"content": "dsfsdfsdf"})
}

//发布
func Publish(c *gin.Context) {
	objId := bson.NewObjectIdWithTime(time.Now())
	DB.Insert("articles", bson.M{
		"_id":              objId,
		"uid":              objId,
		"title":            "test",
		"public_time":      time.Now(),
		"article_classify": "go",
		"article_content":  "content",
		"view_count":       0,
	})
	c.String(http.StatusOK,"ok")
	c.JSON(http.StatusOK,gin.H{
		"_id": objId,
		"ret":true,
	})
}
