package controllers

import (
	. "Todo-Go/models"
)

type TodoController struct {
	BaseController
}

func (this *TodoController) Get() {
	this.TplName = "todo.tpl"
}

func (this *TodoController) Post() {

}

type SearchController struct {
	BaseController
}

func (this *SearchController) Post() {
	content := this.GetString("value")

	todos, _ := SearchTodo(content)

	this.Data["todos"] = todos
	this.Data["js"] = true
	this.TplName = "search.tpl"
}
