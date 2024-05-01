package slack

import (
	"net/url"
	"os"
)

const (
	authorizeURL = "https://slack.com/oauth/v2/authorize"
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
