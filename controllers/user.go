package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"microBlog/DB"
	"gopkg.in/mgo.v2/bson"
	"microBlog/DataManager"
	"time"
	"encoding/json"
	"strings"
)

//登陆
func Login(c *gin.Context) {
	cookie, _ := c.Cookie("session_id")
	//直接登录
	if DataManager.CheckUserExist(cookie) {
		c.Redirect(http.StatusMovedPermanently, "/home")
		return
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		queryRet := DB.Query("user", bson.M{"userName": username, "password": password})

		//用户名或密码正确
		if len(queryRet) > 0 {
			_id := queryRet[0]["_id"]
			value, ok := _id.(bson.ObjectId)
			if ok {
				fmt.Printf("类型匹配ObjectID:%s\n", value.Hex())
			} else {
				fmt.Println("类型不匹配int\n")
				return
			}
			c.SetCookie("session_id", value.Hex(), 604800, "", "localhost", false, true)
			//设置session
			DataManager.SetUserSession(value.Hex(), DataManager.UserSession{value.Hex(), username, password})
			c.Redirect(http.StatusMovedPermanently, "/home")
			return;
		} else {
			//用户或者密码输入有误
			c.HTML(http.StatusOK, "login.html", gin.H{"message": "用户名或者密码输入错误"})
		}
	}
}

//登录界面
func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//注册
func Register(c *gin.Context) {

	//检查用户名是否已经存在
	username := c.PostForm("username")
	password := c.PostForm("password")
	queryRet := DB.Query("user", bson.M{"userName": username})
	//用户名
	if len(queryRet) > 0 {
		c.HTML(http.StatusOK, "register.html", gin.H{"message": "用户名已存在"})
		return
	}
	//插入
	err := DB.Insert("user", bson.M{"userName": username, "password": password, "create": time.Now()})
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"message": "插入失败"})
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)

}

//注册界面
func ToRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"message": ""})
}

//主页
func Home(c *gin.Context) {
	if !CheckLogin(c) {
		return
	}

	userObjId := bson.ObjectIdHex(GetUid(c))

	articleList := DB.Query("articles", bson.M{"uid": userObjId})
	/*处理分类 处理格式如下
		var articleList = {
			"golang":[
				{"title":3333,"id":333}
				{"title":3333,"id":333}
				{"title":3333,"id":333}
			]
		}
	*/
	//进行分类
	tempMap := make(map[string][]map[string]string)
	for _, m := range articleList {
		classify := m["classify"].(string)
		title := m["title"].(string)
		objId := BsonObjIdToString(m["_id"])
		oneArticle := map[string]string{"title": title, "id": objId}
		if _, ok := tempMap[classify]; !ok {
			//不存在
			tempMap[classify] = make([]map[string]string,0)
		}
		tempMap[classify] = append(tempMap[classify],oneArticle)
	}

	typeList := DB.Query("category", bson.M{"uid": userObjId})
	var categoryList []string
	if len(typeList) > 0 {
		categoryStr := typeList[0]["classify"].(string)
		if len(categoryStr) > 0 {
			categoryList = strings.Split(categoryStr, ",")
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"articleList": articleList,
		"typeList":    categoryList,
		"classify":    tempMap,
	})
}

func CheckLogin(c *gin.Context) bool {
	retBool := true
	cookie, _ := c.Cookie("session_id")
	//可以直接登录
	if !DataManager.CheckUserExist(cookie) {
		c.Redirect(302, "/user/toLogin")
		retBool = false
	}
	return retBool
}

func GetUid(c *gin.Context) string {
	var userSessionGet = DataManager.UserSession{}
	sessionid, _ := c.Cookie("session_id")
	ses := DataManager.GetUserSession(sessionid)
	errShal := json.Unmarshal([]byte(ses), &userSessionGet)
	if errShal != nil {
		fmt.Println(errShal)
	}
	return userSessionGet.Id
}

func BsonObjIdToString(objId interface{}) string {
	value, ok := objId.(bson.ObjectId)
	if ok {
		return value.Hex()
	} else {
		return ""
	}
}
