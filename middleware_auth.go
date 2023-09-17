package main

import (
	"fmt"
	"net/http"

	"github.com/kevalsabhani/rss-aggregator/internal/auth"
	"github.com/kevalsabhani/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) middlewareAuth(handler func(http.ResponseWriter, *http.Request, database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
