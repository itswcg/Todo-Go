package controllers

type TodoController struct {
	BaseController
}

func (this *TodoController) Get() {
	this.TplName = "index.tpl"
}

func (this *TodoController) Post() {

}