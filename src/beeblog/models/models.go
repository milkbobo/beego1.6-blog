package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orw:"inde"`
	views           int64     `orm"index"`
	TopicTime       time.Time `orm"index"`
	TopicCount      time.Time
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm"index"`
	Updated         time.Time `orm"index"`
	views           int64     `orm"index"`
	Author          string
	ReplyTime       time.Time `orm"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {

	orm.Debug = true
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/jblog?charset=utf8")
	fmt.Println("已经进入了RegisterDB")

}
