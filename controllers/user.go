package controllers

import (
	"github.com/astaxie/beego"
	. "Todo-Go/models"
)

type RegistorController struct {
	BaseController
}

func (this *RegistorController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if !CheckUserName(username) {
		this.Redirect("/login", 302)
	}


}