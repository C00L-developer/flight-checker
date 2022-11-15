package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// port flag
	flagPort = "port"
	// app name
	appName = "flight-path"
	// version represents the program based on the git tag
	version = "v0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "port",
			Aliases:  []string{"p"},
			Usage:    "gRPC port",
			Required: true,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "runserver",
			Aliases: []string{},
			Usage:   "Run the hermez core",
			Action:  runServer,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}
}
