package todolist

import (
	"fmt"
	"project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 创建待办事项
func Register(c *gin.Context) {
	user := User{}
	err := c.BindJSON(&user)
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

	userModel := models.User{
		Model:    gorm.Model{},
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Status:   user.Status,
		Token:    user.Token,
	}

	// 添加到数据库
	err2 := models.DB.Create(&userModel).Error

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
			Data:  userModel.Username,
			Msg:   "注册成功",
			Error: "",
		})
	}

}
