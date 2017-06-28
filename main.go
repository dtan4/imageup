package main

import (
	"fmt"
	"os"

	"github.com/dtan4/imageup/config"
	"github.com/dtan4/imageup/server"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	server.Run(conf.Port)
}
