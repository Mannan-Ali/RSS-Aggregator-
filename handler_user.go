package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// the main purpose of this function is to take data from the user request and use it to call the database functoin user it acts as a brigde
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	//this are the parameters we get from req body
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

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	//this function now gets called in middlewareAuth with authenticatedUser

	//the whole code here was made into middleware (middleware_auth_handler) so can be used multiple times
	//now we made this functoin  to match the signature of the middlware auth handler (authedHander) by adding userdatabase.user
	// so it can be passed in middlewareAuth function
	responseWithJSON(w, 200, databaseUsertoUser(user))
}
