package main

import (
	"fmt"
	"github.com/urizennnn/GO-PROJECTS/internal/auth"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (Config *apiConfig) middlwareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error %v", err))
			return
		}
		user, err := Config.DB.FetchUsersbyApi(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get User %s", err))
			return
		}
		handler(w, r, user)
	}
}
