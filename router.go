package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("GET /agent", app.handlerFetchAgentData)
	mux.HandleFunc("GET /location", app.handlerSpaceshipLocation)

	return mux
}
