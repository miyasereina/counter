package middleware

import (
	"counter/V2"
	"github.com/gin-gonic/gin"
)

///统计用户数可以直接放在用户登录处做计数，为了方便直接统计请求
func Counter() func(*gin.Context) {
	return func(ctx *gin.Context) {
		weekCt := V2.Init("week")
		monthCt := V2.Init("month")
		weekCt.Incr(1)
		monthCt.Incr(1)
	}

}
