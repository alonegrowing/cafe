package controller

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/web/model"
)

type HomeController interface {
	GetHomePageData(ctx context.Context, id int64) model.HomePageData
}
