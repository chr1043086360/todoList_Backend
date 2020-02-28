package todolist

import (
	"net/http"
	model "project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
)

func ListAll(c *gin.Context) {
	// 从数据库中查询所有代办事项

	var todo []model.TodoList
	err := model.DB.Find(&todo).Error

	if err != nil {
		c.JSON(50001, serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "数据库查询异常",
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  200,
			Data:  todo,
			Msg:   "获取成功",
			Error: "",
		})
	}

	// 返回前端
}

func OneList(c *gin.Context) {
	// 获取参数id

	// 在数据库中查找

	// 返回前端
}
