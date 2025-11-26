package handler

import "github.com/gin-gonic/gin"
import "todo_list_demo/service"

func NewRouter(services ...interface{}) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	{
		// 测试是否连接成功
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  "success form todo_lsit_demo",
			})
		})

		v1.POST("/create", service.Create)

		v1.DELETE("/task/:id", service.DeleteById)

		v1.POST("/task/book", service.BookTask)

		v1.GET("/task/list", service.ListTasks)
	}
	return r
}
