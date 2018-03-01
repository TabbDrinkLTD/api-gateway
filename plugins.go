package main

import (
	"github.com/micro/go-plugins/micro/header"
	"github.com/micro/micro/api"
)

func init() {
	api.Register(header.NewPlugin())
}
