package main

import (
	"github.com/micro/go-plugins/micro/header"
	"github.com/micro/micro/plugin"
)

func init() {
	plugin.Register(header.NewPlugin())
}
