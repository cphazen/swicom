package main

import (
	"example/swift-comply/app"
	"net/http"
)

func main() {
	app := app.Initialize()
	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
