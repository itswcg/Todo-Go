package routers

import (
	"Todo-Go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TodoController{}, "get:GetTodos")
	beego.Router("/add", &controllers.TodoController{}, "post:AddTodo")
	beego.Router("/search", &controllers.TodoController{}, "post:SearchTodo")
	beego.Router("/update/:todo_id", &controllers.TodoController{}, "post:UpdateTodo")
	beego.Router("/del/:todo_id", &controllers.TodoController{}, "get:DelTodo")
	beego.Router("/do/:todo_id", &controllers.TodoController{}, "get:DoTodo")
	beego.Router("/undo/:todo_id", &controllers.TodoController{}, "get:UndoTodo")
	beego.Router("/edit", &controllers.TodoController{}, "get:GetEdit")
	beego.Router("/edit", &controllers.TodoController{}, "post:EditTask")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/setting", &controllers.SettingController{})
}
