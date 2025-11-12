package main

import (
	"fmt"
	"net/http"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/Mannan-Ali/RSS-Aggregator/internal/database/auth"
)

// any func that has this signature can be a authHandler
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// Think of the chi router as a simple switchboard operator.
// The router expects the second argument to be a specific kind of function
// v1Router.Get("/users", someFunction)
// The operator's rule is: "someFunction must be a function that I know how to call. I only know how to call functions that look like this: func(w http.ResponseWriter, r *http.Request)."
// The router doesn't know anything about your database, and it certainly has no idea how to create or pass a database.User object.
//hence we have to return a function here with func(w http.ResponseWriter, r *http.Request) signature , now inside this function we can do othter stuff
//but the signature should match the chi router functions demand

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	//we are returning a clouser
	return func(w http.ResponseWriter, r *http.Request) {

		//now to crate an user the people dont actually need apikeys as we are generationg for them
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
		//here we are calling the handlerGetUserByAPIKey function with authenticated user
		handler(w, r, user)
	}
}
