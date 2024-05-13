package service

import (
	"net"
	"os"

	"slackbot/ent"

	"github.com/slack-go/slack"
)

type Service struct {
	Database    *ent.Client
	SlackClient *slack.Client
}

func (s *Service) dialSocket() (net.Conn, error) {
	c, err := net.Dial("unix", "/tmp/bmo.sock")
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewService(entClient *ent.Client) (Service, error) {
	s := Service{
		Database:    entClient,
		SlackClient: slack.New(os.Getenv("SLACK_TOKEN")),
	}

	return s, nil
}
