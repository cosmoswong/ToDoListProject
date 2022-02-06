package routers

import (
	"ToDoListProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("v1/todo", &controllers.TodoController{})
	beego.Router("v1/todo/:id", &controllers.TodoController{})
}
