package internal

import "time"

func NewApp(config Config, server Server, capturectl CaptureCtl) App {
	return App{
		config: config,
		server: server,
		capturectl: capturectl,
	}
}

type App struct {
	config Config
	server Server
	capturectl CaptureCtl
}

func (a *App) Run() error {
	if a.config.Capture {
		return a.runCapture()
	}
	return a.runServe()
}

func (a *App) runServe() error {
	a.server.Port = a.config.Port

	defer a.server.Close()
	if err := a.server.Serve(); err != nil {
		return err
	}

	return nil
}

func (a *App) runCapture() error {
	go a.runServe()

	time.Sleep(2 * time.Second)

	if err := a.capturectl.Capture(); err != nil {
		return err
	}

	return a.server.Close()
}
