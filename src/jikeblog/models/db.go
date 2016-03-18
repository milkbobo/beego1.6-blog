package models

import (
	"jikeblog/models/class"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/jblog?charset=utf8&loc=Asia%2FShanghai")

	orm.RegisterModel(new(class.User), new(class.Article))
	orm.RunSyncdb("default", false, true)
}
