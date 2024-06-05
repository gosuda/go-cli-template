package app

import "github.com/rs/zerolog"

func NewApp(logger zerolog.Logger, mode string) *App {
	return &App{
		logger: logger,
		mode:   mode,
	}
}

type App struct {
	logger zerolog.Logger
	mode   string
}

func (a *App) Run() error {

	return nil
}

func (a *App) Stop() error {

	return nil
}
