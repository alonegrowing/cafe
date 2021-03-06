package resource

import (
	"github.com/alonegrowing/cafe/pkg/basic/util"
	"github.com/alonegrowing/cafe/pkg/config"
	"github.com/alonegrowing/cafe/pkg/sea/sql"
)

func init() {
	NewMysqlGroup(config.ServiceConfig.Database)
}

func NewMysqlGroup(database []sql.GroupConfig) {
	if len(database) == 0 {
		return
	}
	for _, d := range database {
		g, err := sql.NewGroup(d.Name, d.Master, d.Slaves)
		util.PanicIfError(err)
		err = sql.DefaultGroupManager.Add(d.Name, g)
		util.PanicIfError(err)
	}
	return
}

func Get(service string) *sql.Group {
	return sql.DefaultGroupManager.Get(service)
}
