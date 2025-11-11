package auth

import (
	"errors"
	"net/http"
	"strings"
)

// the main function of this func is to extract api key
// from the headers of an http request
// Example
// Authoriation: APiKey {insert api key here }  we are looking for this format header
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	//why 2 as above specified the format is "APIKey": (actual value of apikey)
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}

// Think of an API key as a long, secret password that a user sends with every request to prove who they are. Unlike JWTs,
//  which are short-lived and contain claims, API keys are typically long-lived (or even permanent until revoked)
//  and are just a unique identifier.
//  You give the user an API key when they sign up. This is their secret credential.
// The user uses that API key to authenticate themselves for all future protected actions ("other tasks").
