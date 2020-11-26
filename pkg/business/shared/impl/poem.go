package impl

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/core/service/impl"

	"github.com/alonegrowing/cafe/pkg/business/shared"
	"github.com/alonegrowing/cafe/pkg/core/model"
	"github.com/alonegrowing/cafe/pkg/core/service"
)

type PoemSharedImpl struct {
	poemService service.PoemService
}

var DefaultPoemShared shared.PoemShared

func init() {
	DefaultPoemShared = NewPoemSharedImpl()
}

func NewPoemSharedImpl() *PoemSharedImpl {
	return &PoemSharedImpl{
		poemService: impl.DefaultPoemService,
	}
}

func (r *PoemSharedImpl) GetPoemByID(ctx context.Context, id int64) *model.Poem {
	return r.poemService.GetPoemByID(ctx, id)
}

func (r *PoemSharedImpl) GetPoemList(ctx context.Context) []*model.Poem {
	return r.poemService.GetPoemList(ctx)
}
