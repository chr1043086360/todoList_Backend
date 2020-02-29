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
	//fmt.Println(user.Username)
	//fmt.Println(user.Password)


	// 判断是否有cookie, 有cookie不用登录
	//ck, _ := c.Cookie("token")
	//if ck == user.Username {
	//	c.JSON(200, serializer.Response{
	//		Code:  200,
	//		Data:  ck,
	//		Msg:   "您已经登录",
	//		Error: "",
	//	})
	//}

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

				c.SetCookie("token", userModel.Username, 300,"/","localhost",false,true)

				c.JSON(200, serializer.Response{
					Code:  200,
					Data:  200,
					Msg:   "登录成功",
					Error: "",
				})
			}else {
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
