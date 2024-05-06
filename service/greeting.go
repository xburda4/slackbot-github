package service

import (
	"slackbot/api/openapi"

	"github.com/slack-go/slack"
)

func (s *Service) greet(command openapi.CommandBody) error {
	_, err := s.SlackClient.OpenView(command.TriggerID, slack.ModalViewRequest{
		Type: "modal",
		Title: &slack.TextBlockObject{
			Type: "plain_text",
			Text: "Greetings",
		},
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.SectionBlock{
					Type: "section",
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: "Random *text*",
					},
				},
			},
		},
		Close:  nil,
		Submit: nil,
	})
	if err != nil {
		return err
	}

	return nil
}
