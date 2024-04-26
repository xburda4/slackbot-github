package slack

import (
	"os"

	"github.com/slack-go/slack"
)

type Service struct {
	client *slack.Client
}

func NewSlackService() Service {
	s := Service{}
	s.client = slack.New(os.Getenv("SLACK_TOKEN"))

	return s
}
