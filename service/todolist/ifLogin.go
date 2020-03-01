package todolist

import "github.com/gin-gonic/gin"

var c gin.Context

// IfLogin 判断是否登录, 如果登录可以拥有一些权限
// ck:cookie(token)   status:1为登录, 2为未登录
func IfLogin() (status int, resCk string ) {
	ck, err := c.Cookie("token")
	if err != nil {
		panic("cookie获取失败")
	}
	if ck == ""{
		return 2, ""
	}else {
		return 1, ck
	}
}
