package controllers

import (
	"Todo-Go/models"
	"github.com/astaxie/beego"
	"strconv"
)

type BaseController struct {
	beego.Controller
	UserId   int64
	UserName string
}

func (this *BaseController) Prepare() {
	sessionid := this.Ctx.GetCookie("sessionid")
	user, err := models.CheckSession(sessionid)
	if err != nil || user == nil {
		this.Redirect("/login", 302)
		return
	} else {
		user_id, _ := strconv.ParseInt(user.(string), 10, 64)
		this.Data["isLogin"] = true
		this.UserId = user_id
		user, _ := models.GetUser(user_id)
		this.UserName = user.(string)
	}
}
