package github_classroom

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

const BaseClassroomApiURL = "https://api.github.com"

type ClassroomClient struct {
	BaseURL    string
	httpClient *http.Client
}

// NewClient creates a new ClassroomClient with the given access token
func NewClient(ctx context.Context, accessToken string) *ClassroomClient {

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})

	httpClient := oauth2.NewClient(ctx, ts)
	return &ClassroomClient{
		BaseURL:    BaseClassroomApiURL,
		httpClient: httpClient,
	}
}
