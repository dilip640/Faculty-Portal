package main

import (
	"os"

	"github.com/dilip640/Faculty-Portal/server"
	"github.com/dilip640/Faculty-Portal/storage"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Faculty Portal"
	app.Usage = "A Faculty Portal at Academic University"
	app.Author = "Jainam | Dilip"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "runserver",
			Aliases: []string{"s"},
			Usage:   "Starts the server",
			Action: func(c *cli.Context) {
				srv := server.NewInstance()
				srv.Start()
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Database migration",
			Action: func(c *cli.Context) {
				storage.Migrate()
			},
		},
		{
			Name:    "downonestep",
			Aliases: []string{"ds1"},
			Usage:   "Database roll back",
			Action: func(c *cli.Context) {
				storage.DownOneStep()
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
