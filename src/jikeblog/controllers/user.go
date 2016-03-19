package controllers

//	"github.com/astaxie/beego"

import (
	"jikeblog/models/class"
)

type UserController struct {
	//	beego.Controller
	BaseController
}

func (c *UserController) Profile() {

	//	c.Data["userid"] = "geek"
	//	c.Data["tag"] = "I am a geek"
	//	c.Data["hobby"] = []string{"chess", "code"}

	//获取用户信息
	id := c.Ctx.Input.Param(":id")
	u := &class.User{Id: id}
	u.ReadDB()

	c.Data["u"] = u

	//获取所有文章列表
	a := &class.Article{Author: u}
	as := a.Gets()

	c.Data["articles"] = as

	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplName = "user/setting.html"
}
