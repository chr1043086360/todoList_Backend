package todolist

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project2019/models"
	"project2019/serializer"
)

//var flag int

func Change(c *gin.Context){
	// 根据id获取对象
	id := c.Params.ByName("id")

	// 获取属性
	//status := c.PostForm("status")
	todo := models.TodoList{}
	err := c.BindJSON(&todo)

	if err != nil {
		c.JSON(400, serializer.Response{
			Code:  400,
			Data:  nil,
			Msg:   "请求错误",
			Error: err.Error(),
		})
		fmt.Println(err)
	}

	err2:= models.DB.Model(&todo).Where("id = ?", id).Update("status",todo.Status).Error

	// 返回状态码
	if err2 != nil{
		c.JSON(400, serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "修改失败",
			Error: err2.Error(),
		})
	}else {
		c.JSON(200, serializer.Response{
			Code:  200,
			Data:  todo.Title,
			Msg:   "修改成功",
			Error: "",
		})
	}
}