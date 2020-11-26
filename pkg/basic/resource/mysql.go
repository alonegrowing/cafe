package resource

import (
	"github.com/alonegrowing/cafe/pkg/config"
	"github.com/alonegrowing/cafe/pkg/sea/sql"
)

var DefaultDB *sql.Group

func init() {
	_ = NewMysqlGroup(config.ServiceConfig.Database)
}

func NewMysqlGroup(database []sql.SQLGroupConfig) error {
	if len(database) == 0 {
		return nil
	}
	for _, d := range database {
		g, err := sql.NewGroup(d.Name, d.Master, d.Slaves)
		if err != nil {
			return err
		}
		err = sql.SQLGroupManager.Add(d.Name, g)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(service string) *sql.Group {
	return sql.SQLGroupManager.Get(service)
}
