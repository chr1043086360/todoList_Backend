package todolist

import (
	"fmt"
	"project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 创建待办事项
func CreateTodoList(c *gin.Context) {
	todo := Todo{}
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
	// 绑定模型

	todoListModel := models.TodoList{
		Model:  gorm.Model{},
		Id:     todo.Id,
		Title:  todo.Title,
		Status: todo.Status,
		Info:   todo.Info,
	}

	// 添加到数据库
	err2 := models.DB.Create(&todoListModel).Error

	if err2 != nil {
		c.JSON(50001, serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "保存失败",
			Error: err2.Error(),
		})
		fmt.Println(err2)
	} else {
		c.JSON(200, serializer.Response{
			Code:  200,
			Data:  serializer.BuildTodo(todoListModel),
			Msg:   "",
			Error: "",
		})
	}

}
