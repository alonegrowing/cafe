package main

import (
	"fmt"

	"github.com/alonegrowing/cafe/pkg/config"
	"github.com/alonegrowing/cafe/pkg/web"
)

func main() {
	route := web.NewRouter()
	_ = route.Run(fmt.Sprintf(":%d", config.ServiceConfig.Service.WEBPort))
}
