package handler

import (
	"strconv"

	"github.com/alonegrowing/cafe/pkg/web/controller"
	"github.com/alonegrowing/cafe/pkg/web/controller/impl"
	"github.com/gin-gonic/gin"
)

type HomePageHandler struct {
	homeController controller.HomeController
}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{
		homeController: impl.DefaultHomeController,
	}
}

func (m *HomePageHandler) Get(r *gin.Context) {
	id, _ := strconv.ParseInt(r.Request.FormValue("id"), 10, 64)
	homePageData := m.homeController.GetHomePageData(r, id)
	r.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    homePageData,
	})
}
