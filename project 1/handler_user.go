package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	user,err:=Config.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:params.Name,
	})
	if err != nil {
		respondWithError(w,400,"Error Createing User")
		return
	}
	respondWithJSON(w, 200, user)
}
