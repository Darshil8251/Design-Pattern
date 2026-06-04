package http

import (
	"net/http"
)

type httpClient struct {
	// require
	client *http.Client

	// Optional
	auth      string
	rateLimit int
}

type optional func(h *httpClient)

func CreateHTTPClient(options ...optional) *httpClient {
	client := &httpClient{
		client: &http.Client{},
	}

	// Apply optional settings
	for _, option := range options {
		option(client)
	}
	return client

}

func WithAuth(auth string) optional {
	return func(h *httpClient) {
		h.auth = auth
	}
}

func WithRateLimit(rateLimit int) optional {
	return func(h *httpClient) {
		h.rateLimit = rateLimit
	}
}
