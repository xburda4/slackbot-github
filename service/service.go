package service

import (
	"os"

	"slackbot/ent"

	"github.com/slack-go/slack"
)

type Service struct {
	Database    *ent.Client
	SlackClient *slack.Client
}

func NewService(entClient *ent.Client) (Service, error) {
	s := Service{
		Database:    entClient,
		SlackClient: slack.New(os.Getenv("SLACK_TOKEN")),
	}

	return s, nil
}
