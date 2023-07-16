package app

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
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
}
