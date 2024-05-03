package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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

type GithubUserDocument struct {
	SlackID           string `json:"slack_id"`
	GithubUsername    string `json:"github_username"`
	GithubAccessToken string `json:"github_access_token"`
}

func (s *Service) GithubOauth(ctx context.Context, code, encodedState string) error {
	//decodedState contains SlackID
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
	ghUser, _, err := ghClient.WithAuthToken(githubOauth.AccessToken).Users.Get(context.Background(), "")
	if err != nil {
		return err
	}

	now := time.Now()
	err = s.Database.
		GithubUser.Create().
		SetID(uuid.New()).
		SetSlackID(string(decodedState)).
		SetGhUsername(*ghUser.Login).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
