package routers

import (
	"jikeblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("api/user/profile", &controllers.UserController{}, `get:Profile`)
}
