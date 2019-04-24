package controllers

import (
	. "Todo-Go/models"
	"Todo-Go/utils"
	"github.com/astaxie/beego"
	"strconv"
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

	user_id, err := AddUser(username, password)
	if err != nil {
		this.Redirect("/register", 302)
		return
	}

	sessionid := utils.GenerateSession(user_id)

	_, err = AddSession(sessionid, user_id)
	if err != nil {
		this.Redirect("/register", 302)
		return
	}

	this.Ctx.SetCookie("sessionid", sessionid, 2592000)

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

	user, login := CheckPassword(username, password)

	if !login {
		this.Redirect("/login", 302)
		return
	}

	user_id, _ := strconv.ParseInt(user.(string), 10, 64)

	sessionid := utils.GenerateSession(user_id)

	_ = UpdateSession(sessionid)

	this.Ctx.SetCookie("sessionid", sessionid, 2592000)

	this.Redirect("/", 302)
}

type SettingController struct {
	BaseController
}

func (this *SettingController) Get() {
	this.TplName = "setting.tpl"
}

func (this *SettingController) Post() {
	username := this.UserName
	old_password := this.GetString("old_password")
	new_password := this.GetString(("new_password"))
	confirm_password := this.GetString("confirm_password")

	_, check := CheckPassword(username, old_password)

	if new_password != confirm_password || check == false {
		this.Redirect("setting", 302)
		return
	}

	_ = ChangePassword(username, new_password)

	this.Redirect("/", 302)
}

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	this.Ctx.SetCookie("sessionid", "")
	this.Redirect("/", 302)
}
