package handler

import (
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	gin.HandlersChain
}

func (m *BaseHandler) render(r *gin.Context, data interface{}) {
	r.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}
