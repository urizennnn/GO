package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

func (Config *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error Parsing JSON")
		return
	}
	feed, err := Config.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error Creating User", err))
		return
	}

	respondWithJSON(w, 201, databasetoFeed(feed))
}
