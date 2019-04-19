package controllers

import (
	. "Todo-Go/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.tpl"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	confirm_password := this.GetString("confirm_password")

	if password != confirm_password || !CheckUserName(username) {
		this.Redirect("/register", 302)
		return
	}

	_, err := AddUser(username, password)
	if err != nil {
		this.Redirect("/register", 302)
		return
	}

	this.Ctx.SetCookie("sessionid", "")

	this.Redirect("/", 302)
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	login := CheckPassword(username, password)

	if !login {
		this.Redirect("/login", 302)
	}

	// login

	this.Redirect("/", 302)
}
