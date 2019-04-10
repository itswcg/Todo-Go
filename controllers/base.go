package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	sessionid := this.GetSession("sessionid")
	if sessionid == nil {
		this.Redirect("/login", 302)
	} else {
		// 验证session 返回user_id
	}
}
