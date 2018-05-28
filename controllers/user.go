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
			//1，向浏览器设置session id
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "session_id",
				Value:    value.String(),
				MaxAge:   604800,
				HttpOnly: true,
			})
			fmt.Println("------",value.Hex())
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
func ToRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"message": ""})
}

func Home(c *gin.Context) {
	CheckLogin(c)
	var userSessionGet = DataManager.UserSession{}
	sessionid, _ := c.Cookie("session_id")
	ses := DataManager.GetUserSession(sessionid)
	errShal := json.Unmarshal([]byte(ses), &userSessionGet)
	if errShal != nil {
		fmt.Println(errShal)
	}
	//进行分类
	var tempArray []string
	tempMap :=make(map[string]int)
	articleList := DB.Query("articles", bson.M{"uid": bson.ObjectIdHex(userSessionGet.Id)})
	//处理分类
	for index,m:=range articleList{
		classify :=m["classify"].(string)

		if _, ok := tempMap[classify]; ok {
			continue
		} else {
			tempMap[classify] = index
			tempArray = append(tempArray, classify)
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articleList": articleList,
		"typeList":    tempArray,
	})
}

func CheckLogin(c *gin.Context) {
	cookie, _ := c.Cookie("session_id")
	//可以直接登录
	if !DataManager.CheckUserExist(cookie) {
		c.Redirect(302, "/user/toLogin")
	}
}
