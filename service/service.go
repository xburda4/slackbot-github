package service

import "slackbot/service/slack"

type Service struct {
	SlackService slack.Service
}

func NewService() (Service, error) {
	s := Service{}
	s.SlackService = slack.NewSlackService()

	return s, nil
}
