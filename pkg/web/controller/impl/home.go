package impl

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/web/controller"

	"github.com/alonegrowing/cafe/pkg/business/shared/impl"

	"github.com/alonegrowing/cafe/pkg/business/shared"

	"github.com/alonegrowing/cafe/pkg/web/model"
)

type HomeControllerImpl struct {
	poemShared shared.PoemShared
}

var DefaultHomeController controller.HomeController

func init() {
	DefaultHomeController = NewHomeControllerImpl()
}

func NewHomeControllerImpl() *HomeControllerImpl {
	return &HomeControllerImpl{
		poemShared: impl.DefaultPoemShared,
	}
}

func (r *HomeControllerImpl) GetHomePageData(ctx context.Context, id int64) model.HomePageData {
	poem := r.poemShared.GetPoemByID(ctx, id)
	return model.HomePageData{
		Id:        poem.Id,
		Token:     poem.Token,
		Title:     poem.Title,
		Author:    poem.Author,
		Intro:     poem.Intro,
		Imgs:      poem.Imgs,
		Content:   poem.Content,
		UpdatedAt: poem.UpdatedAt,
		CreatedAt: poem.CreatedAt,
	}
}
