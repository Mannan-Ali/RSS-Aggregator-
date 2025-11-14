package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollower(w http.ResponseWriter, r *http.Request, user database.User) {
	//this are the parameters we get from req body
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	//parsing the req body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollower(r.Context(), database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feed follower: %v", err))
		return
	}
	responseWithJSON(w, 201, databaseFeedFollowertoFeedFollower(feedFollow))
}

func (apiCfg *apiConfig) handlerGetAllFollowersFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	allUserFeeds, err := apiCfg.DB.GetAllFollowersFeeds(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudn't get users feeds : %v", err))
		return
	}
	responseWithJSON(w, 201, databaseUserFeedstoUserFeeds(allUserFeeds))
}
func (apiCfg *apiConfig) handlerUnfollowUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFllowIdStr := chi.URLParam(r, "feedFollowID")
	feedIdToUnfollow, err := uuid.Parse(feedFllowIdStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudn't parse feed id to be unfollowed : %v", err))
	}
	err = apiCfg.DB.UnfollowUserFeed(r.Context(), database.UnfollowUserFeedParams{
		ID:     feedIdToUnfollow,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudn't unfollow feed : %v", err))
		return
	}
	responseWithJSON(w, 200, struct{}{})
}
