package controllers

import (
	"github.com/astaxie/beego"
	. "Todo-Go/models"
)

type TodoController struct {
	BaseController
}

func (this *TodoController) Get() {
	this.TplName = "index.tpl"
}

func (this *TodoController) Post() {

}