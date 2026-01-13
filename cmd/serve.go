package cmd

import (
	"Todo-list/config"
	"Todo-list/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
