package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"slackbot/api/openapi"

	"entgo.io/ent/dialect/sql"
	"github.com/google/go-github/v61/github"
	"github.com/slack-go/slack"
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
	decodedState, err := base64.StdEncoding.DecodeString(encodedState)
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
		SetSlackID(string(decodedState)).
		SetGhUsername(*ghUser.Login).
		SetGhAccessToken(githubOauth.AccessToken).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		OnConflictColumns("slack_id").
		UpdateGhAccessToken().
		UpdateUpdatedAt().
		UpdateGhUsername().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

const (
	authorizeURL       = "https://slack.com/oauth/v2/authorize"
	githubAuthorizeURL = "https://github.com/login/oauth/authorize"
)

func (s *Service) Authenticate() (string, error) {
	parsedURL, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}

	q := parsedURL.Query()
	q.Set("scope", os.Getenv("BOT_SCOPES"))
	q.Set("client_id", os.Getenv("SLACK_CLIENT_ID"))
	q.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
	parsedURL.RawQuery = q.Encode()

	return parsedURL.RequestURI(), nil
}

func (s *Service) Authorize() error {

	//slack.GetBotOAuthToken(http.Client{}, os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGNING_SECRET"))
	return nil
}

func (s *Service) githubLogin(_ context.Context, command openapi.CommandBody) error {
	encodedState := base64.StdEncoding.EncodeToString([]byte(command.UserID))

	_, err := s.SlackClient.PostEphemeral(command.ChannelID, command.UserID, slack.MsgOptionBlocks(slack.ActionBlock{
		Type: "actions",
		Elements: &slack.BlockElements{
			ElementSet: []slack.BlockElement{
				slack.ButtonBlockElement{
					Type: "button",
					Text: &slack.TextBlockObject{
						Type: "plain_text",
						Text: "Login",
					},
					URL: fmt.Sprintf("%s?client_id=%s&redirect_uri=%s",
						githubAuthorizeURL,
						os.Getenv("GITHUB_CLIENT_ID"),
						fmt.Sprintf("%s&state=%s", os.Getenv("GITHUB_REDIRECT_URI"), encodedState)),
				},
			},
		},
	}))
	if err != nil {
		return nil
	}

	return nil
}

func (s *Service) githubLogout(ctx context.Context, command openapi.CommandBody) error {
	_, err := s.Database.GithubUser.Delete().Where(sql.FieldContains("slack_id", command.UserID)).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = s.SlackClient.PostEphemeral(command.ChannelID, command.UserID, slack.MsgOptionText("You were successfully logged out of Github.", false))
	if err != nil {
		return err
	}

	return nil
}
