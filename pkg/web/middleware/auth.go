package middleware

import (
	"github.com/alonegrowing/cafe/pkg/basic/macro"

	"github.com/gin-gonic/gin"
)

// 登陆校验 中间件
func Auth() gin.HandlerFunc {
	return func(r *gin.Context) {
		id := r.Query("id")
		if id == "2" {
			r.JSON(200, gin.H{
				"code":    macro.StatusAuthFailed,
				"message": macro.ErrMsg[macro.StatusAuthFailed],
				"data":    []int{},
			})
			r.Abort()
		}
		r.Next()
	}
}
