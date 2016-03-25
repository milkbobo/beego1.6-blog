package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}
