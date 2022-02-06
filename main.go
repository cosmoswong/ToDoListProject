package main

import (
	"ToDoListProject/models"
	_ "ToDoListProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.InitMysql()
	beego.Run()
}

