package controller

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/web/model"
)

type HomeController interface {
	GetPoemById(ctx context.Context, id int64) *model.Poem
	GetPoemList(ctx context.Context) []*model.Poem
}
