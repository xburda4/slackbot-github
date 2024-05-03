package slack

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"

	"slackbot/api/openapi"

	"github.com/slack-go/slack"
)

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

func (s *Service) githubLogin(command openapi.CommandBody) error {
	var encodedState []byte
	base64.StdEncoding.Encode(encodedState, []byte(command.UserID))

	_, err := s.client.PostEphemeral(command.ChannelID, command.UserID, slack.MsgOptionBlocks(slack.ActionBlock{
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
