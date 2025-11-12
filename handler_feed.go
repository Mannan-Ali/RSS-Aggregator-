package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	//this are the parameters we get from req body
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	//parsing the req body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}
	responseWithJSON(w, 201, databaseFeedtoFeed(feed))
}
