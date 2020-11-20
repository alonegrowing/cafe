package shared

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/core/model"
)

type PoemShared interface {
	GetPoemByID(ctx context.Context, id int64) model.Poem
}
