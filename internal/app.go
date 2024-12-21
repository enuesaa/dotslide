package internal

func NewApp(server Server) App {
	return App{
		server: server,
	}
}

type App struct {
	Port int
	Workdir string
	
	server Server
}

func (a *App) Run() error {
	a.server.Port = a.Port

	defer a.server.Close()
	if err := a.server.Serve(); err != nil {
		return err
	}

	return nil
}
