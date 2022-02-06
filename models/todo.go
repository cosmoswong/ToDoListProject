package models
// Todo Model
type Todo struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Status bool `json:"status"`
}

func CerateATodo(todo *Todo) error {
   return DB.Create(todo).Error
}

func GetATodo(id string) (todo *Todo,err error) {
    if err =DB.Where("id = ?",id).First(&todo).Error;err != nil{
        return nil,err
    }
    return
}
func GetAllTodo() (todos []*Todo,err error)  {
    if err = DB.Find(&todos).Error;err !=nil{
        return nil,err
    }
    return 
}
func UpdateATodo(todo *Todo) error  {
    return DB.Model(todo).Updates(map[string]interface{}{
        "title":todo.Title,
        "status":todo.Status,
    }).Error
}
func DeleteATodo(id string) error  {
    return DB.Where("id =?",id).Delete(&Todo{}).Error
}
func (Todo) TableName() string {
    return "Todo"
}