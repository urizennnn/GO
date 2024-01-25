package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

func (Config *apiConfig) handlerCreateFeedfollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error Parsing JSON")
		return
	}
	followFeed, err := Config.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error Creating User", err))
		return
	}

	respondWithJSON(w, 201, databaseFollowsFeed(followFeed))
}

func (Config *apiConfig) GetUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := Config.DB.GetUserFeed(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get feed follows %v", err))
		return
	}

	respondWithJSON(w, 200, UserFeed(feedFollow))
}
func (Config *apiConfig) DeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIDStr := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Deleting Feed %v", err))
		return
	}
	err = Config.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		ID:     feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Deleting Feed %v", err))
		return
	}
	respondWithJSON(w,200,"Successful")
}
