package api

import "net/http"

type API struct {
	PAT    string
	client *http.Client
	BaseURI string
}

func Client(pat string) *API {
	return &API{
		client: &http.Client{},
		PAT:    pat,
		BaseURI: "https://gloapi.gitkraken.com/v1/glo",
	}
}
