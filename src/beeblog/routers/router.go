package routers

import (
	"beeblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, `get:Get`)
	beego.Router("/category", &controllers.CategoryController{}, `get:Get;post:Post`)
	beego.Router("/login", &controllers.LoginController{}, `get:Get;post:Post`)

}
