package service

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/core/model"
)

type PoemService interface {
	GetPoemByID(ctx context.Context, id int64) model.Poem
}
