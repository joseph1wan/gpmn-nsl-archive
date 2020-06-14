package main

func setupTestApp() *App {
	app := App{
		conf: ServerConfig{Port: 4000},
	}
	app.setupRoutes()

	//TODO: add setting up a test postgresql database
	return &app
}
