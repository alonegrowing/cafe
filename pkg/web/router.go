package web

import (
	"io/ioutil"

	"github.com/alonegrowing/cafe/pkg/web/handler"
	"github.com/alonegrowing/cafe/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	route := gin.Default()
	route.Use(middleware.Logger())

	route.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	homeHandler := handler.NewHomePageHandler()

	route.GET("/api/poem", homeHandler.Get)
	route.GET("/api/poems", homeHandler.GetList)

	return route
}
