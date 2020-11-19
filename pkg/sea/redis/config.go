package redis

import (
	"fmt"
)

type Config struct {
	ServerName     string `json:"server_name"`
	Addr           string `json:"addr"`
	Password       string `json:"password"`
	MaxIdle        int    `json:"max_idle"`
	MaxActive      int    `json:"max_active"`
	IdleTimeout    int    `json:"idle_timeout"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	Database       int    `json:"database"`
	SlowTime       int    `json:"slow_time"`
	Retry          int    `json:"retry"`
}

func (o *Config) init() error {
	if o.ServerName == "" {
		return fmt.Errorf("redis: ServerName not allowed empty string")
	}
	if o.Addr == "" {
		return fmt.Errorf("redis: Addr not allowed empty string")
	}
	if o.Database < 0 {
		return fmt.Errorf("redis: Database less than zero")
	}

	if o.MaxIdle < 0 {
		o.MaxIdle = 100
	}
	if o.MaxActive < 0 {
		o.MaxActive = 100
	}
	if o.IdleTimeout < 0 {
		o.IdleTimeout = 100
	}
	if o.ReadTimeout < 0 {
		o.ReadTimeout = 50
	}
	if o.WriteTimeout < 0 {
		o.WriteTimeout = 50
	}
	if o.ConnectTimeout < 0 {
		o.ConnectTimeout = 300
	}

	if o.SlowTime <= 0 {
		o.SlowTime = 100
	}

	if o.Retry < 0 {
		o.Retry = 0
	}

	return nil
}
