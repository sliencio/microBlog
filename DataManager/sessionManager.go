package DataManager

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/gin-gonic/gin/json"
)

type UserSession struct {
	Id       string `json:"session_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var conn redis.Conn

func init() {
	//Connect
	c, err := redis.Dial("tcp", "localhost:6379")
	conn  = c
	if err != nil {
		panic(err)
	}
	response, err := conn.Do("AUTH", "lien")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

}

func SetUserSession(key string, userS UserSession) error{
	value, err := json.Marshal(userS)
	if err != nil {
		panic(err)
		return err
	}
	_, err2 := conn.Do("PUSH", key, value,"EX",3600)
	if err2 != nil {
		panic(err2)
		return err2
	}
	return err2
}

func GetUserSession(key string) string {
	v, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return v
}

func CheckUserExist(key string)bool{
	is_key_exit, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}
	return is_key_exit
}
