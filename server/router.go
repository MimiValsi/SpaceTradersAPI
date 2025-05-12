package server

import "net/http"

func router() http.Handler {
	mux := http.NewServeMux()

	return mux
}
