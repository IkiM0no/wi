package main

import (
	"os"

	"github.com/urfave/cli"
)

var (
	interval int
	pp       bool
)

func main() {
	app := buildApp()
	app.Run(os.Args)
}

func buildApp() *cli.App {
	app := cli.NewApp()
	app.Name    = "wi"
	app.Version = "0.1.1"
	app.Usage   = "Fetch metrics of WiFi neighbors."
	app.Action  = func(c *cli.Context) error {
		runner()
		return nil
	}
	app.Flags = []cli.Flag{
	        cli.IntFlag{
	                Name:  "i",
	                Usage: "-i <WiFi polling interval in seconds>",
	                Destination: &interval,
	        },
	        cli.BoolFlag{
	                Name:  "p",
	                Usage: "-p <true|false>",
	                Destination: &pp,
	        },
	}
	return app
}
