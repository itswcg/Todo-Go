package controllers

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

}