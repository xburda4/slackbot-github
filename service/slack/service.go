package slack

import (
	"os"

	"slackbot/ent"

	"github.com/slack-go/slack"
)

type Service struct {
	client   *slack.Client
	Database *ent.Client
}

func NewSlackService(entClient *ent.Client) Service {
	s := Service{
		Database: entClient,
	}
	s.client = slack.New(os.Getenv("SLACK_TOKEN"))

	return s
}
