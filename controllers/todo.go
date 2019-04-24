package controllers

import (
	. "Todo-Go/models"
	"strconv"
)

type TodoController struct {
	BaseController
}

func (this *TodoController) GetTodos() {
	user_id := this.UserId
	task, _ := GetTask(user_id)

	if task == nil {
		this.Data["task"] = ""
	} else {
		this.Data["task"] = task[0]["content"]
	}

	todoList_Today, _ := GetTodo(user_id, false, true)
	doList_Today, _ := GetTodo(user_id, true, true)
	todoList, _ := GetTodo(user_id, false, false)
	doList, _ := GetTodo(user_id, true, false)

	this.Data["todoList_Today"] = todoList_Today
	this.Data["doList_Today"] = doList_Today
	this.Data["todoList"] = todoList
	this.Data["doList"] = doList

	this.TplName = "todo.tpl"
}

func (this *TodoController) AddTodo() {
	user_id := this.UserId
	content := this.GetString("content")

	_, err := AddTodo(user_id, content)
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/", 302)
}

func (this *TodoController) DelTodo() {
	todo := this.Ctx.Input.Param(":todo_id")
	todo_id, _ := strconv.ParseInt(todo, 10, 64)
	_ = DeleteTodo(todo_id)

	this.Redirect("/edit", 302)
}

func (this *TodoController) UpdateTodo() {
	todo := this.Ctx.Input.Param(":todo_id")
	content := this.GetString("content")
	todo_id, _ := strconv.ParseInt(todo, 10, 64)
	_ = UpdateTodo(todo_id, content)

	this.Redirect("/edit", 302)
}

func (this *TodoController) DoTodo() {
	todo := this.Ctx.Input.Param(":todo_id")
	todo_id, _ := strconv.ParseInt(todo, 10, 64)
	_ = DoTodo(todo_id, 1)

	this.Redirect("/", 302)
}

func (this *TodoController) UndoTodo() {
	todo := this.Ctx.Input.Param(":todo_id")
	todo_id, _ := strconv.ParseInt(todo, 10, 64)
	_ = DoTodo(todo_id, 0)

	this.Redirect("/", 302)
}

func (this *TodoController) EditTask() {
	user_id := this.UserId
	content := this.GetString("content")

	_, err := AddTask(user_id, content)
	if err != nil {
		this.Redirect("/edit", 302)
	}

	this.Redirect("/edit", 302)
}

func (this *TodoController) GetEdit() {
	user_id := this.UserId
	task, _ := GetTask(user_id)

	if task == nil {
		this.Data["task"] = ""
	} else {
		this.Data["task"] = task[0]["content"]
	}

	todos, _ := GetAllTodo(user_id)

	this.Data["todoList"] = todos
	this.Data["todoJs"] = true
	this.TplName = "edit.tpl"
}

func (this *TodoController) SearchTodo() {
	content := this.GetString("value")

	todos, _ := SearchTodo(content)

	this.Data["todos"] = todos
	this.Data["todoJs"] = true
	this.TplName = "search.tpl"
}
