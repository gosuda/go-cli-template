package main

import (
	"os"

	"github.com/urfave/cli"
)

var (
	// example flag
	mode string
	// input additional flag variable here
)

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "mode",                      // flag name
			Value:       getEnv("MODE", "debug"),     // from env or default value
			Usage:       "run mode (debug, release)", // flag description
			Destination: &mode,                       // variable to store the flag value
		},
		// input additional flag here
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
