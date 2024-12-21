package internal

func NewApp(config Config, server Server) App {
	return App{
		config: config,
		server: server,
	}
}

type App struct {
	config Config
	server Server
}

func (a *App) Run() error {
	a.server.Port = a.config.Port

	defer a.server.Close()
	if err := a.server.Serve(); err != nil {
		return err
	}

	return nil
}
