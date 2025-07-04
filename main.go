package main

import (
	"fmt"

	"github.com/smartbot/auth/api"
	"github.com/smartbot/auth/pkg/config"
)

func main() {
	config.LoadConfig()
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))
}
