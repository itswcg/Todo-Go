package controllers

import (
	. "Todo-Go/models"
)

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.TplName = "register.tpl"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if !CheckUserName(username) {
		this.Redirect("/login", 302)
	}

	_, err := AddUser(username, password)
	if err != nil {
		this.Redirect("/login", 302)
	}

	//设置session login

	this.Redirect("/", 302)
}

type LoginController struct {
	BaseController
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
