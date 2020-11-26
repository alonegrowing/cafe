package handler

import (
	"strconv"

	"github.com/alonegrowing/cafe/pkg/basic/macro"

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
		d.error(r, macro.ErrorParamIllegal)
	}
	poem := d.homeController.GetPoemById(r, id)
	d.success(r, poem)
}

func (d *PoemHandler) GetList(r *gin.Context) {
	poems := d.homeController.GetPoemList(r)
	d.success(r, poems)
}
