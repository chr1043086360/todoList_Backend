package todolist

import (
	"github.com/gin-gonic/gin"
	"project2019/serializer"
)

func Logout(c *gin.Context){

	ck, err := c.Cookie("token")
	
	if err != nil {
		c.JSON(200, serializer.Response{
			Code:  403,
			Data:  403,
			Msg:   "cookie为空",
			Error: "",
		})
	}
	if ck != "" {
		c.SetCookie("token", "", -1,"/","122.51.107.26",false,true)
		c.JSON(200, serializer.Response{
			Code:  200,
			Data:  200,
			Msg:   "登出成功",
			Error: "",
		})
	}
}
