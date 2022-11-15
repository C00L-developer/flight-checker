package main

import (
	"os"
	"os/signal"

	"github.com/C00L-developer/flight-checker/pkg/server"
	"github.com/urfave/cli/v2"
)

func runServer(ctx *cli.Context) error {
	err := server.RunServer(server.Config{
		GRPCPort: "9090",
		HTTPPort: ctx.String(flagPort),
	})

	if err != nil {
		return err
	}

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}
