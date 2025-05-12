package server

import (
	"log"
	"net/http"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/api"
)

func handlerAgentData(client *api.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := client.MyAgent(); err != nil {
			log.Printf("%w", err)
		}

		next.ServeHTTP(w, r)
	})
}
