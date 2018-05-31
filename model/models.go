package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Article struct {
	Id        bson.ObjectId `bson:"_id"`
	Uid       bson.ObjectId `bson:"uid"`
	Title     string        `bson:"title"`
	Time      time.Time     `bson:"time"`
	Classify  string        `bson:"classify"`
	Content   string        `bson:"content"`
	Viewed    int64         `bson:"viewed"`
	Commented int64         `bson:"commented"`
	Theme	  string		`bson:"theme"`
}
type Comment struct {
	Id          bson.ObjectId `bson:"_id"`
	ArticleId   bson.ObjectId `bson:"articleId"`
	Commentator string        `bson:"commentator"`
	Content     string        `bson:"content"`
	Time        time.Time     `bson:"time"`
}

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Email    bson.ObjectId `bson:"email"`
	Password string        `bson:"password"`
	Name     string        `bson:"name"`
	Create   string        `bson:"create"`
}
