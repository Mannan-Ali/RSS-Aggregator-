package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/Mannan-Ali/RSS-Aggregator/internal/database/auth"
	"github.com/google/uuid"
)

// the main purpose of this function is to take data from the user request and use it to call the database functoin user it acts as a brigde
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	//parsing the req body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	// apiCfg.DB gives you access to the collection of all your pre-defined and type-safe database functions.
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}
	responseWithJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	//now to create an user the people dont actually need apikeys as we are generationg for them
	//but to getUsers or your own data will require an api key
	//we always use package name as prefix before calling a function we a different package is useds
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}
	responseWithJSON(w, 200, databaseUsertoUser(user))
}
