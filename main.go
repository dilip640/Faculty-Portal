package main

import (
	"github.com/dilip640/Faculty-Portal/server"
)

func main() {
	srv := server.NewInstance()
	srv.Start()
}
