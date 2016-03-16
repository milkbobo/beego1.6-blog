package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {

	c.Data["userid"] = "geek"
	c.Data["tag"] = "I am a geek"
	c.Data["hobby"] = []string{"chess", "code"}

	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}
