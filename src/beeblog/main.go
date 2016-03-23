package main

import (
	_ "beeblog/routers"

	"beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 启动 beego
	beego.Run()

}
