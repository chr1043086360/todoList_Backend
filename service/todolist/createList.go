package todolist

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
)

// 创建待办事项
func CreateTodoList(c *gin.Context) {

	todo := Todo{}
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(400, serializer.Response{
			Code:  400,
			Data:  nil,
			Msg:   "参数请求错误",
			Error: err.Error(),
		})
		fmt.Println(err)
	} else {
		// 获取cookie里的token用于标识文章所属用户
		token, err4 := c.Cookie("token")
		if err4 != nil {
			c.JSON(200, serializer.Response{
				Code:  200,
				Data:  403,
				Msg:   "无权创建列表, 请先登录",
				Error: "403 Forbidden",
			})
		} else {

			// 在用户表中通过token获取用户
			userModel := models.User{}

			err5 := models.DB.Where("token = ?", token).First(&userModel).Error
			if err5 != nil {
				c.JSON(500, serializer.Response{
					Code:  50001,
					Data:  50001,
					Msg:   "token获取用户异常",
					Error: err5.Error(),
				})
			} else {
				// 拿到用户的用户名添加到refer中
				todoListModel := models.TodoList{
					Model:  gorm.Model{},
					Title:  todo.Title,
					Status: todo.Status,
					Info:   todo.Info,
					Refer:  userModel.Username,
				}


				// 添加到数据库
				err3 := models.DB.Create(&todoListModel).Error

				if err3 != nil {
					c.JSON(50001, serializer.Response{
						Code:  50001,
						Data:  nil,
						Msg:   "保存失败",
						Error: err3.Error(),
					})
					fmt.Println(err3)
				}

				c.JSON(200, serializer.Response{
					Code:  200,
					Data:  serializer.BuildTodo(todoListModel),
					Msg:   "",
					Error: "",
				})
			}


		}
	}

}
