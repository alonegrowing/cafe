package config

import (
	"os"

	"github.com/alonegrowing/cafe/pkg/basic/resource"
	"github.com/alonegrowing/cafe/pkg/sea/log"
	"github.com/alonegrowing/cafe/pkg/sea/redis"
	"github.com/alonegrowing/cafe/pkg/sea/sql"
	"github.com/alonegrowing/cafe/pkg/sea/tomlconfig"
)

var (
	ServiceConfig Config
)

func init() {
	var conf = "./conf/prod/config.toml"
	_ = tomlconfig.ParseTomlConfig(conf, &ServiceConfig)
	InitLoggerConfig(ServiceConfig.Logger)
	_ = resource.NewRedis(ServiceConfig.Redis)
	_ = resource.NewMysqlGroup(ServiceConfig.Database)
}

type Service struct {
	WEBPort int64 `toml:"web_port"`
	RPCPort int64 `toml:"rpc_port"`
}

type Config struct {
	ServiceName string               `toml:"service_name"`
	Service     Service              `toml:"service"`
	Logger      LoggerConfig         `toml:"log"`
	Database    []sql.SQLGroupConfig `toml:"database"`
	Redis       []redis.Config       `toml:"redis"`
}

type LoggerConfig struct {
	Level   log.Level `toml:"level"`
	LogPath string    `toml:"logpath"`
}

func InitLoggerConfig(conf LoggerConfig) {
	log.SetLevel(conf.Level)

	file, _ := os.OpenFile(conf.LogPath, os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
}
