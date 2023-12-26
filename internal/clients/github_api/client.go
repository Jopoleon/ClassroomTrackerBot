package github_api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

const githubAppInstallationsURL = "https://api.github.com/app/installations"

// GetInstallationAccessToken retrieves the installation access token
func GetInstallationAccessToken(installationID int64, jwtToken string) (string, error) {
	url := fmt.Sprintf(githubAppInstallationsURL+"/%v/access_tokens", installationID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	req.Header.Set("Accept", "application/vnd.github.machine-man-preview+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	type tokenResponse struct {
		Token string `json:"token"`
	}

	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}

	return tr.Token, nil
}

func MakeClassroomRequest(ctx context.Context, token string) ([]byte, error) {
	// Example Classroom API endpoint
	url := "https://api.github.com/classrooms"

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)

	// OAuth2 client
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	}))

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	return io.ReadAll(resp.Body)
}
