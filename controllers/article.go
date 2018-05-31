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
	if !CheckLogin(c) {
		return
	}
	typeList := DB.Query("category", bson.M{"uid": bson.ObjectIdHex(GetUid(c))})
	var categoryStr string
	if len(typeList) > 0 {
		categoryStr = typeList[0]["classify"].(string)
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"typeList": categoryStr,
	})
}

//展示文章列表
func Show(c *gin.Context) {
	if !CheckLogin(c) {
		return
	}
	objIDStr := c.Param("articleId")
	objID := bson.ObjectIdHex(strings.Split(objIDStr, "\"")[1])
	fmt.Println("--------", objID)
	//return
	typeList := DB.Query("articles", bson.M{"_id": objID})
	if len(typeList) > 0 {
		article := typeList[0]
		viewCount := article["viewed"].(int64)
		//更新阅读人数
		DB.Update("articles", bson.M{"_id": objID}, bson.M{"viewed": viewCount + 1})
		c.HTML(http.StatusOK, "showArticle.html", gin.H{"content": article})
	}
}
//帮助
func Example(c *gin.Context) {
	c.HTML(http.StatusOK, "exampleTemplate.html", nil)
}

//删除
func Delete(c *gin.Context) {
	objIDStr := c.Param("articleId")
	objID := bson.ObjectIdHex(objIDStr)
	typeList := DB.Query("articles", bson.M{"_id": objID})
	if len(typeList) > 0 {
		//更新阅读人数
		DB.Remove("articles", bson.M{"_id": objID})
	}
}

//编辑
func ReEdit(c *gin.Context) {
	//查询文章内容
	objIDStr := c.Param("articleId")
	objID := bson.ObjectIdHex(strings.Split(objIDStr, "\"")[1])
	articleList := DB.Query("articles", bson.M{"_id": objID})

	typeList := DB.Query("category", bson.M{"uid": bson.ObjectIdHex(GetUid(c))})
	if len(typeList) > 0 && len(articleList)>0 {
		categoryStr := typeList[0]["classify"].(string)
		c.HTML(http.StatusOK, "reEdit.html", gin.H{
			"article":articleList[0],
			"typeList":categoryStr,
		})
	}
}

//发布
func Publish(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	classify := c.PostForm("classify")
	theme := c.PostForm("theme")

	userObjID := bson.ObjectIdHex(GetUid(c))
	//检查该类型是否为新类型。如果是新类型，就加上，如果不是，就不管
	checkClassifyIsNew(userObjID,classify)

	article := &model.Article{
		Id		:bson.NewObjectId(),
		Uid		:userObjID,
		Title	:title,
		Time	:time.Now(),
		Classify:classify,
		Content	:content,
		Theme	:theme,
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

//发布
func RePublish(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	classify := c.PostForm("classify")
	articleIdStr := c.PostForm("articleId")
	theme := c.PostForm("theme")
	userObjID := bson.ObjectIdHex(GetUid(c))

	articleId := bson.ObjectIdHex(articleIdStr)

	//检查该类型是否为新类型。如果是新类型，就加上，如果不是，就不管
	checkClassifyIsNew(userObjID,classify)

	//进行更新操作
	err := DB.Update("articles", bson.M{"_id":articleId},bson.M{"title":title,"time":time.Now(),"classify":classify,"content":content,"theme":theme})
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

func checkClassifyIsNew(uid bson.ObjectId,classify string){
	typeList := DB.Query("category", bson.M{"uid": uid})
	var categoryList []string
	var categoryStr string
	//用户已经有分类
	if len(typeList) > 0 {
		categoryStr = typeList[0]["classify"].(string)
		categoryList = strings.Split(categoryStr, ",")
		var isHadClassify = false
		for _, value := range categoryList {
			if value == classify {
				isHadClassify = true
			}
		}

		//新类型，进行插入操作
		if !isHadClassify {
			var retInsertStr = ""
			if len(categoryStr) > 0 {
				retInsertStr = categoryStr + "," + classify
			} else {
				retInsertStr = classify
			}
			DB.Update("category", bson.M{"uid": uid}, bson.M{"classify": retInsertStr})
		}

	} else {
		//用户尚未设置分类
		DB.Insert("category", bson.M{"uid": uid, "classify": classify})
	}
}
