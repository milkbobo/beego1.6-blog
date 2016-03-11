package main

import (
	_ "jikeblog/models"
	"jikeblog/models/class"
	_ "jikeblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	class.TestORM()
	beego.Run()
}
