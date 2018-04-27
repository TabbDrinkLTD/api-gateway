package main

import (
	auth "github.com/TabbDrinkLTD/auth-service/micro"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
)

func init() {
	plugin.Register(cors.NewPlugin())
	plugin.Register(auth.NewPlugin())
}
