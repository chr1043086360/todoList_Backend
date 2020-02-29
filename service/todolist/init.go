/*
  Author ： CHR_崔贺然
  Time ： 2020
  TODO ： init文件为包下所有服务创建Bind的结构体
*/
package todolist

type Todo struct {
	Id     uint   `json:"id"`
	Title  string `json:"title" form:"title" binding:"required,max=20"`
	Status bool   `json:"status" form:"status"`
	Info   string `json:"info" form:"info"`
}

type User struct {
	Id     uint   `json:"id"`
	Username  string `json:"username" form:"username" binding:"required,max=20"`
	Password string   `json:"password" form:"password" binding:"required,max=20"`
	Token   string `json:"token" form:"token"`
	Nickname string `json:"nickname" form:"nickname"`
	Status string `json:"status" form:"status"`
}
