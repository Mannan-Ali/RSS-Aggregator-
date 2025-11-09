package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// her empty interface is passed, that means anything can be of this interface type
// like we can pass map,slice,arr anything as data
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//json.Marshal is the function that translates your Go data into JSON text.
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshel JSON response")
		w.WriteHeader(500)
		return
	}
	//adding header to response
	// 1.	w.WriteHeader(code) sets the Status Line - there is only 1 status field
	w.Header().Add("Content-Type", "application/json")
	// 	2.	w.Header().Add(...) adds a Header Field.
	// These are key-value pairs that provide additional information or metadata about the response.
	// You can have many header fields.
	w.WriteHeader(code)
	w.Write(dat)

}

// here we have string instead of interface as with error we will only pass what is the error
func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	//`json:"error"` this is a key and we are telling the marshal function to match it with the json error key when passed
	// 	The json.Marshal function would turn the Go string "Something went wrong" into a JSON string "Something went wrong".
	// The raw HTTP response body would look like this: "Something went wrong"
	type errResponse struct {
		Error string `json:"error"`
	}
	//  with adding struct
	// 	{
	//     "error": "Something went wrong"
	// }
	responseWithJSON(w, code, errResponse{
		Error: msg,
	})
}
