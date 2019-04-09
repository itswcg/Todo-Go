package controllers

import (
	"github.com/astaxie/beego"
	. "Todo-Go/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	todos, err := GetAllTodo(1)
	if err == nil {
		c.Data["todos"] = todos
	}
	c.Data["content"] = "hello"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

}