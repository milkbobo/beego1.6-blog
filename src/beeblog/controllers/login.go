package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "ok"
	fmt.Println(isExit)
	if isExit {
		this.Ctx.SetCookie("uname", "", 0, "/")
		this.Ctx.SetCookie("pwd", "", 0, "/")
		fmt.Println("isExit")
		this.Redirect("/", 301)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	//	this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//	return

	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	fmt.Println(uname)
	fmt.Println(pwd)
	fmt.Println(autoLogin)

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 1000
		if autoLogin {
			maxAge = 1<<31 - 1

		}
		fmt.Println("setCookie")
		fmt.Println(uname)
		fmt.Println(pwd)
		fmt.Println(maxAge)
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")

	}

	this.Redirect("/", 301)
	return

}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		fmt.Println(err)
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		fmt.Println("false pwd")
		return false
	}
	pwd := ck.Value
	fmt.Println("uname pwd ok!")
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}
