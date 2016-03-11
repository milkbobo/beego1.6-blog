package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"

	this.Data["TrueCond"] = true
	this.Data["FalseCond"] = false

	type u struct {
		Name string
		Age  int
		Sex  string
	}

	user := &u{
		Name: "Joe",
		Age:  20,
		Sex:  "Male",
	}

	this.Data["User"] = user

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	this.Data["Nums"] = nums

	this.Data["TplVal"] = "hey guys"

	this.Data["Html"] = "<div>hello beego</div>"

	this.Data["Pipe"] = "<div>hello beego</div>"

}
