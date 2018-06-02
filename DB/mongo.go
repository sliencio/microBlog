package DB

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"microBlog/config"
	"time"
)

var operator Operater

type Operater struct {
	mogSession *mgo.Session
	database   *mgo.Database
	dbUrl      string
}

func init() {
	var dburl = config.MongoUrl
	var dbName = config.MongoDbName

	fmt.Println("start Connect mongo: url:" + dburl + " database:" + dbName)
	operator = Operater{nil, nil, dburl}
	operator.connect()
}

//连接数据库
func (operater *Operater) connect() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"127.0.0.1"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  config.MongoDbName,
		Source:    "admin",
		Username:  config.MongoUsername,
		Password:  config.MongoPassword,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	mogsession, err :=mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	operater.mogSession = mogsession
	operater.database = mogsession.DB(config.MongoDbName)

}

//插入
func InsertByMap(collection string, data map[string]string) error {
	collcetion := operator.database.C(collection)
	err := collcetion.Insert(data)
	return err
}

//插入
func Insert(collection string, seletor interface{}) error {
	collcetion := operator.database.C(collection)
	err := collcetion.Insert(seletor)
	return err
}

//更新一行
func Update(collection string, oldData interface{}, upateData interface{}) error {
	collcetion := operator.database.C(collection)
	err := collcetion.Update(oldData, bson.M{"$set": upateData})
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//更新多行
func UpdateAll(collection string, oldData interface{}, upateData interface{}) error {
	collcetion := operator.database.C(collection)
	_, err := collcetion.UpdateAll(oldData, bson.M{"$set": upateData})
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//统计集合中数据的个数
func Query(collection string, seletor interface{}) (retData []map[string]interface{}) {
	collcetion := operator.database.C(collection)
	query := collcetion.Find(seletor)
	ps := []map[string]interface{}{}
	query.All(&ps)
	return ps
}

func QueryBind(collection string, seletor interface{},bind interface{}){
	collcetion := operator.database.C(collection)
	err := collcetion.Find(seletor).All(bind)
	if err!=nil{
		panic(err)
		return
	}
}

//单行删除
func Remove(collection string, seletor interface{}) error {
	collcetion := operator.database.C(collection)
	return collcetion.Remove(seletor)
}

func RemoveAll(collection string, seletor interface{}) error {
	collcetion := operator.database.C(collection)
	_, err := collcetion.RemoveAll(seletor)
	return err
}
