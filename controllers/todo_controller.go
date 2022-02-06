package controllers

import (
    "ToDoListProject/models"
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
)

type TodoController struct {
    beego.Controller
}

func (this *TodoController) Get()  {

    todos, err := models.GetAllTodo()
    if err!=nil{
        this.responseErr(err)
        return
    }else {
        this.Data["json"]=map[string]interface{}{"code": http.StatusOK,"message":todos}
    }
    this.ServeJSON()

}

func (this *TodoController) Post()  {
    data := this.Ctx.Input.RequestBody
    todo := new(models.Todo)
    err := json.Unmarshal(data, todo)
    if err!=nil{
        this.responseErr(err)
        return
    }
    err = models.CerateATodo(todo)
    if err != nil{
        this.responseErr(err)
        return
    }else {
        this.Data["json"]=map[string]interface{}{"code": http.StatusOK,"message":"success"}
    }
    this.ServeJSON()
}
func (this *TodoController) Put()  {
    id:=this.Ctx.Input.Param(":id")

    aTodo, err := models.GetATodo(id)
    if err != nil{
        this.responseErr(err)
        return
    }

    data := this.Ctx.Input.RequestBody
    err = json.Unmarshal(data, aTodo)
    if err!=nil{
        this.responseErr(err)
        return
    }

    err = models.UpdateATodo(aTodo)
    if err != nil{
        this.responseErr(err)
        return
    }else {
        this.Data["json"]=map[string]interface{}{"code": http.StatusOK,"message":"success"}
    }
    this.ServeJSON()
}
func (this *TodoController) Delete()  {
    id:=this.Ctx.Input.Param(":id")

     err := models.DeleteATodo(id)
    if err != nil{
        this.responseErr(err)
        return
    }else {
        this.Data["json"]=map[string]interface{}{"code": http.StatusOK,"message":"success"}
    }
    this.ServeJSON()
}

func (this *TodoController) responseErr(err error) {
    this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
    this.ServeJSON()
}
