package service

import (
	"context"

	slackapi "github.com/slack-go/slack"
)

func (s *Service) ProcessReceivedSlackMessage(_ context.Context, msg *slackapi.Msg) error {
	_, _, err := s.SlackClient.PostMessage(msg.Channel, slackapi.MsgOptionText(msg.Text, false))
	return err
}
