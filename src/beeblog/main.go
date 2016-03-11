package main

import (
	_ "beeblog/routers"
	"fmt"

	"beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
	fmt.Println("models.RegisterDB()")
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true

	fmt.Println("Debug")
	// 自动建表
	orm.RunSyncdb("default", false, true)
	fmt.Println("RunSyncdb~")
	// 启动 beego
	beego.Run()

}
