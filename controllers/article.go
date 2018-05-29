package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"microBlog/DB"
	"microBlog/model"
	"net/http"
	"time"
	"strings"
)

//编辑
func Edit(c *gin.Context) {
	if !CheckLogin(c){
		return
	}
	typeList:=DB.Query("category",bson.M{"uid":bson.ObjectIdHex(GetUid(c))})
	var categoryStr string
	if len(typeList)>0{
		categoryStr = typeList[0]["classify"].(string)
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"typeList":categoryStr,
	})
}

//展示文章列表
func Show(c *gin.Context) {
	if !CheckLogin(c){
		return
	}
	objID := c.Param("articleId")
	typeList:=DB.Query("articles",bson.M{"_id":bson.ObjectIdHex(objID)})
	if len(typeList)>0{
		c.HTML(http.StatusOK, "showArticle.html", gin.H{"content": typeList[0]})
	}
}

//发布
func Publish(c *gin.Context) {
	content:=c.PostForm("content")
	title:=c.PostForm("title")
	classify:=c.PostForm("classify")
	fmt.Println("-----uid----:",GetUid(c))

	userObjId := bson.ObjectIdHex(GetUid(c))

	typeList:=DB.Query("category",bson.M{"uid":userObjId})
	var categoryList []string
	var categoryStr string
	if len(typeList)>0{
		categoryStr = typeList[0]["classify"].(string)
		categoryList=strings.Split(categoryStr,",")
	}
	var isHadClassify = false
	for _,value :=range categoryList{
		if value == title{
			isHadClassify = true
		}
	}

	//新类型，进行插入操作
	if !isHadClassify{
		var retInsertStr = ""
		if len(categoryStr)>0{
			retInsertStr=retInsertStr+","+title
		}else{
			retInsertStr = title
		}
		DB.Insert("category",bson.M{"uid":userObjId,"classify": retInsertStr})
	}


	article := &model.Article{
		Id:       bson.NewObjectId(),
		Uid:      userObjId,
		Title:    title,
		Time:     time.Now(),
		Classify: classify,
		Content:  content,
	}
	err := DB.Insert("articles", article)
	fmt.Println("err:", err)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": true,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"ret": false,
		})
	}
}
