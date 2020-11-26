package handler

import (
	"strconv"

	"github.com/alonegrowing/cafe/pkg/web/controller"
	"github.com/alonegrowing/cafe/pkg/web/controller/impl"
	"github.com/gin-gonic/gin"
)

type PoemHandler struct {
	BaseHandler
	homeController controller.HomeController
}

func NewHomePageHandler() *PoemHandler {
	return &PoemHandler{
		homeController: impl.DefaultHomeController,
	}
}

func (d *PoemHandler) Get(r *gin.Context) {
	id, err := strconv.ParseInt(r.Request.FormValue("id"), 10, 64)
	if id <= 0 || err != nil {
		d.render(r, 1)
	}
	poem := d.homeController.GetPoemById(r, id)
	d.render(r, poem)
}

func (d *PoemHandler) GetList(r *gin.Context) {
	poems := d.homeController.GetPoemList(r)
	d.render(r, poems)
}
