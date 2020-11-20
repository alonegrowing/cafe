package impl

import (
	"context"

	"github.com/alonegrowing/cafe/pkg/basic/dao"
	"github.com/alonegrowing/cafe/pkg/basic/macro"
	"github.com/alonegrowing/cafe/pkg/basic/util"
	"github.com/alonegrowing/cafe/pkg/sea/sql"
)

type PoemDaoImpl struct {
	tableName string
	group     *sql.Group
}

var DefaultPoemDao dao.PoemDao

func init() {
	DefaultPoemDao = NewPoemDaoImpl()
}

func NewPoemDaoImpl() *PoemDaoImpl {
	return &PoemDaoImpl{
		tableName: "poem",
		group:     sql.Get(macro.DBName),
	}
}

func (r *PoemDaoImpl) GetPoemById(ctx context.Context, id int64) (poem *dao.Poem) {
	poem = &dao.Poem{}
	if err := r.group.Slave().Table(r.tableName).Where("id=?", id).Limit(1).Scan(&poem).Error; err != nil {
		util.PanicIfError(err)
	}
	return
}
