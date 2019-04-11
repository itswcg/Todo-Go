package routers

import (
	"Todo-Go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TodoController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
}
