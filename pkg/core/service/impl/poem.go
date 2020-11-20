package impl

import (
	"context"
	"strings"

	"github.com/alonegrowing/cafe/pkg/basic/dao"
	"github.com/alonegrowing/cafe/pkg/basic/dao/impl"
	"github.com/alonegrowing/cafe/pkg/core/model"
	"github.com/alonegrowing/cafe/pkg/core/service"
)

type PoemServiceImpl struct {
	poemDao dao.PoemDao
}

var DefaultPoemService service.PoemService

func init() {
	DefaultPoemService = NewPoemServiceImpl()
}

func NewPoemServiceImpl() *PoemServiceImpl {
	return &PoemServiceImpl{
		poemDao: impl.DefaultPoemDao,
	}
}

func (r *PoemServiceImpl) GetPoemByID(ctx context.Context, id int64) model.Poem {
	poem := r.poemDao.GetPoemById(ctx, id)
	return model.Poem{
		Id:        poem.Id,
		Token:     poem.Token,
		Author:    poem.Author,
		Title:     poem.Title,
		Intro:     poem.Intro,
		Imgs:      strings.Split(poem.Imgs, ","),
		Content:   poem.Content,
		Status:    poem.Status,
		UpdatedAt: poem.UpdatedAt,
		CreatedAt: poem.CreatedAt,
	}
}
