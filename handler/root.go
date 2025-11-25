package handler

import "github.com/gin-gonic/gin"

func NewRouter(service ...interface{}) *gin.Engine {
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
	}
	return r
}
