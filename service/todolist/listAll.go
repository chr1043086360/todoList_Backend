package todolist

import (
	model "project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
)

// ListAll 默认获取我的todoList,如果有用户登录则获取该用户列表
func ListAll(c *gin.Context) {

	var todo []model.TodoList
	userModel := model.User{}

	// 获取登录状态拿到用户
	token, err9 := c.Cookie("token")
	if err9 != nil {
		// 如果没有cookie默认获取我的todoList

		err8 := model.DB.Where("refer = ?", "chr").Find(&todo).Error
		if err8 != nil {
			c.JSON(500, serializer.Response{
				Code:  50001,
				Data:  50001,
				Msg:   "数据库获取失败",
				Error: err8.Error(),
			})
		}else {
			// 返回我的列表
			c.JSON(200, serializer.Response{
				Code:  200,
				Data:  todo,
				Msg:   "OK",
				Error: "",
			})
		}

	}else {
		// 有cookie获取到该用户的
		// 根据用户token获取用户
		err7 := model.DB.Where("token = ?", token).First(&userModel).Error
		if err7 != nil {
			c.JSON(500, serializer.Response{
				Code:  50001,
				Data:  50001,
				Msg:   "根据用户token获取用户失败",
				Error: err7.Error(),
			})

		}else {
			// 返回用户发的todoList

			err6 := model.DB.Where("refer = ?", userModel.Username).Find(&todo).Error
			if err6 != nil{
				c.JSON(500, serializer.Response{
					Code:  50001,
					Data:  50001,
					Msg:   "返回用户发的todoList失败",
					Error: err6.Error(),
				})
			}

			c.JSON(200, serializer.Response{
				Code:  200,
				Data:  todo,
				Msg:   "OK",
				Error: "",
			})
		}
	}

	// 从数据库中查询所有代办事项
	// 没有用户登录权限的版本v1


	//err := model.DB.Find(&todo).Error
	//
	//if err != nil {
	//	c.JSON(50001, serializer.Response{
	//		Code:  50001,
	//		Data:  nil,
	//		Msg:   "数据库查询异常",
	//		Error: err.Error(),
	//	})
	//} else {
	//	c.JSON(http.StatusOK, serializer.Response{
	//		Code:  200,
	//		Data:  todo,
	//		Msg:   "获取成功",
	//		Error: "",
	//	})
	//}

}

func OneList(c *gin.Context) {
	// 获取参数id

	// 在数据库中查找

	// 返回前端
}
