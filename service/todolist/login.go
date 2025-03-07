package todolist

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project2019/models"
	"project2019/serializer"
)

//var flag int

func Login(c *gin.Context){

	// 获取属性
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

	var userModel models.User
	// 在数据库中修改
	//fuck := models.DB.Where("id = ?", id).First(&user)
	err2 := models.DB.Where("username = ?", user.Username).First(&userModel).Error

	if err2 != nil {
		c.JSON(200, serializer.Response{
			Code:  200,
			Data:  50001,
			Msg:   "没有此用户",
			Error: err2.Error(),
		})
	} else{
		if user.Username == userModel.Username {
			if user.Password == userModel.Password{

				// 将用户名加上服务器id存在用户表的token中

				err9 := models.DB.Model(&userModel).Where("username = ?", user.Username).Update("token",user.Username+"666").Error

				if err9 != nil {
					c.JSON(500, serializer.Response{
						Code:  50001,
						Data:  50001,
						Msg:   "更新token错误",
						Error: err9.Error(),
					})
				}else{
				// 设置缓存用于登录
				c.SetCookie("token", userModel.Token, 500000,"/","122.51.107.26",false,true)


				c.JSON(200, serializer.Response{
					Code:  200,
					Data:  200,
					Msg:   "登录成功",
					Error: "",
				})
				}
			} else {
				c.JSON(200, serializer.Response{
					Code:  200,
					Data:  40001,
					Msg:   "账号密码错误",
					Error: "",
				})
			}
		}
	}

}
