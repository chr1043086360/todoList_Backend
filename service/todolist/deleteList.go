package todolist

import (
	"fmt"
	"net/http"
	"project2019/models"
	"project2019/serializer"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	// 获取前端id参数
	id := c.Params.ByName("id")
	fmt.Println(id)
	// 到数据库中查找删除
	todoModel := models.TodoList{}

	// 这里必须是Error
	err := models.DB.Where("id = ?", id).Delete(&todoModel).Error

	if err != nil {
		c.JSON(50001, serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "数据库查询失败",
			Error: err.Error(),
		})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  200,
			Data:  nil,
			Msg:   "删除陈成功",
			Error: "",
		})
	}
	// 返回状态
}
