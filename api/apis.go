/*
  Author ： CHR_崔贺然
  Time ： 2019.11.14
  Description ： 所有的接口都在这里定义，还有中间件的启动,注意顺序不能改变
*/

package api

import (

	// "io"

	// "time"

	// "log"
	"io"
	"net/http"
	"os"
	"project2019/middleware"
	"project2019/service"
	"project2019/service/todolist"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

// Api is 路由接口
func Api() *gin.Engine {

	// 初始化一个引擎(包含Logger，Recovery中间件)
	r := gin.Default()

	// !生产模式配置
	gin.SetMode(gin.ReleaseMode)

	// *无中间件启动
	// r := gin.New()

	// *自定义HTTP配置

	// s := &http.Server{
	// 	Addr:           ":6666",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	// *一行代码配置Https
	// log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	// 中间件
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	// *禁用控制台颜色
	// gin.DisableConsoleColor()

	// *创建记录日志的文件
	f, _ := os.Create("gin.log")
	// 标准输出
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 设置静态文件访问路径
	r.Static("/static", "./static")
	r.StaticFS("./more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// r.LoadHTMLFiles("static/form.html")

	r.LoadHTMLGlob("template/*")

	// 定义一组接口
	//v1 := r.Group("/api/v1")
	{
		// ping接口用来做生死检查
		// v1.GET("/ping", service.Ping)

		// 登陆接口
		// v1.POST("/login", service.Login)

		// 获取路径参数

		// v1.GET("/param/:name", service.Param)

		//post接口获取列表数据

		// v1.POST("/articlelist", service.ArticleList)

		// 只绑定get数据
		// v1.Any("/testing", service.FormDemo)

		// 重定向
		// v1.GET("/redirect", service.Redirect)

	}

	// * 测试接口
	v2 := r.Group("/api/v2")
	{
		// 测试redis get数据接口
		v2.GET("/redisTest", service.RedisGetKey)

		// 测试redis连接池
		v2.GET("/redisPool", service.RedisPool)

		// 上传文件到七牛云
		v2.GET("/qiniu", service.QiniuTest)
	}

	// * 项目api
	/*******************************************************************************
	*                                 Golang 大法好!                                *
	*******************************************************************************/
	v3 := r.Group("/api/v3")
	{
		// 测试接口
		v3.GET("/", todolist.Index)

		// 注册接口
		v3.POST("/register", todolist.Register)

		// 登录接口
		v3.POST("/login", todolist.Login)

		// 登出接口
		v3.GET("/logout", todolist.Logout)

		// 添加待办事项
		v3.POST("/todo", todolist.CreateTodoList)

		// 查询所有待办事项
		v3.GET("/todo", todolist.ListAll)

		// 查询某一个待办事项
		v3.GET("/todo/:id", todolist.OneList)

		// 修改待办事项
		v3.PUT("/todo/:id", todolist.Change)

		// 删除待办事项
		v3.DELETE("/todo/:id", todolist.Delete)

		// 外键通过username查询文章
		v3.GET("/fk", todolist.FkSearch)
	}

	return r

}

// Gin 项目下可以运行多个服务：https://www.kancloud.cn/shuangdeyu/gin_book/949444

// 优雅的shutdown：https://www.kancloud.cn/shuangdeyu/gin_book/949445

// 使用go-asserts将服务器构建成一个包含模板的二进制文件：https://www.kancloud.cn/shuangdeyu/gin_book/949446

// 使用自定义的结构绑定表单：https://www.kancloud.cn/shuangdeyu/gin_book/949447

// 将请求踢绑定到不同的结构体中：https://www.kancloud.cn/shuangdeyu/gin_book/949448

// HTTP/2 服务器推送：https://www.kancloud.cn/shuangdeyu/gin_book/949449
