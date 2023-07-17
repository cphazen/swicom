package main

import (
	"example/swift-comply/app"
	"example/swift-comply/database"
	"net/http"
)

func main() {
	app := app.Initialize()
	app.DB = &database.DB{}
	err := app.DB.Open()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
