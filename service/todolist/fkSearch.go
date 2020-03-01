package todolist

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model "project2019/models"
	"project2019/serializer"
)


// FkSearch 外键查询练习
func FkSearch(c *gin.Context){
	//  前端传来的数据绑定结构体用于获取
	//err := c.BindJSON(&user)
	usernameRef := c.Query("username")
	// 如果绑定失败, 抛错
	//if err != nil {
	//	c.JSON(400, serializer.Response{
	//		Code:  400,
	//		Data:  400,
	//		Msg:   "参数错误,请根据接口文档检查传递的参数",
	//		Error: err.Error(),
	//	})
	//}

	// 通过username获取文章(todoList)列表

	//fmt.Println(usernameRef)
	todoModel := model.TodoList{}
	userModel := model.User{}

	// 绑定userModel
	model.DB.Where("username=?", usernameRef).First(&userModel)

	// 关联代码
	model.DB.Model(&todoModel).Related(&todoModel.User, "Username")

	// 查询
	//result := model.DB.First(&todoModel)

	// 先打个日志看看
	//fmt.Println(result)
	fmt.Println(todoModel.User.Username)
	//fmt.Println(userModel.Username)

	// 创建查询列表容器
	var todoAllByUsername []model.TodoList
	err2 := model.DB.Find(&todoAllByUsername).Error
	if err2 != nil {
		c.JSON(400, serializer.Response{
			Code:  400,
			Data:  50001,
			Msg:   "查询错误",
			Error: err2.Error(),
		})
	}else {
		c.JSON(200, serializer.Response{
			Code:  200,
			Data:  todoAllByUsername,
			Msg:   "OK",
			Error: "",
		})
	}

}
