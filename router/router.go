package router

import (
	"counter/V2"
	"counter/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Counter())
	// ping 测试
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"m": V2.Cts.GetWithIndex("month"),
		})
		return
	})

	api := r.Group("/api")
	api.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})
	return r
}
