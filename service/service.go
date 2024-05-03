package service

import (
	"slackbot/ent"
	"slackbot/service/slack"
)

type Service struct {
	SlackService slack.Service
	Database     *ent.Client
}

func NewService(entClient *ent.Client) (Service, error) {
	s := Service{
		Database:     entClient,
		SlackService: slack.NewSlackService(),
	}

	return s, nil
}
