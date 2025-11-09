package main

import "net/http"

//this function acts as a health check thats why r is not used
// and in interface we pass empty sturct that accounts as empty json
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
