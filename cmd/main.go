package main

import (
	"fmt"
	"os"

	"github.com/gosuda/go-cli-template/app"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:  "Input your cli App Name",
		Usage: "Input your cli App Usage",
		Flags: flags(),
		Action: func(c *cli.Context) error {
			return Run()
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func Run() error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// make your app here
	app := app.NewApp(logger, mode)
	err := app.Run()
	if err != nil {
		logger.Error().Err(err).Msg("error")
		return err
	}

	// Wait to stop
	interrupt := handleKillSig(func() {
		app.Stop()
	}, logger)

	<-interrupt.C
	return nil
}
