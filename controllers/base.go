package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Login bool
	User  int64
}

func (this *BaseController) Prepare() {
	sessionid := this.Ctx.GetCookie("sessionid")
	if sessionid == "" {
		this.Redirect("/login", 302)
	} else {
		// 验证session 返回user_id
	}
}
