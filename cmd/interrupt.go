package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
)

type interrupt struct {
	C chan struct{}
}

func handleKillSig(handler func(), logger zerolog.Logger) interrupt {
	i := interrupt{
		C: make(chan struct{}),
	}

	sigChannel := make(chan os.Signal, 1)

	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for signal := range sigChannel {
			logger.Info().Msgf("Receive signal %s, Shutting down...", signal)
			handler()
			close(i.C)
		}
	}()
	return i
}
