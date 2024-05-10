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
	Socket      *net.Conn
}

func NewService(entClient *ent.Client) (Service, error) {
	c, err := net.Dial("unix", "/tmp/bmo.sock")
	if err != nil {
		return Service{}, err
	}

	s := Service{
		Database:    entClient,
		SlackClient: slack.New(os.Getenv("SLACK_TOKEN")),
		Socket:      &c,
	}

	return s, nil
}
