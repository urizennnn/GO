package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/urizennnn/GO-PROJECTS/internal/auth"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

func (Config *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error Parsing JSON")
		return
	}
	user, err := Config.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error Creating User",err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (Config *apiConfig) getUserbyApi(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, 200, databaseUserToUser(user))

}
