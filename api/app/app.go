package app

import (
	"example/swift-comply/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *database.DB
}

func Initialize() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.routes()
	return a
}

func (a *App) routes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/cats", a.GetCatsHandler()).Methods("GET")
}
