package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
	this.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") +
		"\nhttpport:" + beego.AppConfig.String("httpport") +
		"\nrunmode:" + beego.AppConfig.String("runmode"))

	hp := strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	this.Ctx.WriteString("\n\nappname:" + beego.BConfig.AppName +
		"\nhttpport:" + hp +
		"\nrunmode:" + beego.BConfig.RunMode)

	beego.Trace("test1")
	beego.Info("test2")

	beego.SetLevel(beego.LevelInformational)

	beego.Trace("test233")
	beego.Info("test144")
}
