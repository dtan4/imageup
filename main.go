package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dtan4/imageup/server"
)

const (
	defaultPort = 8000
)

func main() {
	var port int

	if os.Getenv("PORT") == "" {
		port = defaultPort
	} else {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		port = p
	}

	server.Run(port)
}
