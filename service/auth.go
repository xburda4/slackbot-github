package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v61/github"
	"github.com/google/uuid"
)

const (
	githubOauthURL = "https://github.com/login/oauth/access_token"
)

type GithubOauthResp struct {
	AccessToken string `json:"access_token" form:"access_token"`
	Scope       string `json:"scope" form:"scope"`
	TokenType   string `json:"token_type" form:"token_type"`
	ExpiresIn   int    `json:"expires_in" form:"expires_in"`
}

func (s *Service) GithubOauth(code, encodedState string) error {
	var decodedState []byte
	_, err := base64.StdEncoding.Decode(decodedState, []byte(encodedState))
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s", githubOauthURL, code, os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_CLIENT_SECRET")), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	var githubOauth GithubOauthResp
	if err := json.NewDecoder(resp.Body).Decode(&githubOauth); err != nil {
		return err
	}

	ghClient := github.Client{}
	user, _, err := ghClient.WithAuthToken(githubOauth.AccessToken).Users.Get(context.Background(), "")
	if err != nil {
		return err
	}

	err = s.Database.Github.Create().
		SetSlackID(string(decodedState)).
		SetID(uuid.New()).
		SetUsername(*user.Login).
		SetToken(githubOauth.AccessToken).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
