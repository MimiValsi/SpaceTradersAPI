package main

import "net/http"

func (c *apiCfg) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("GET /agent/data", c.handlerAgent)

	return mux
}
