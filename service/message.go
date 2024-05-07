package service

import (
	"context"

	slackapi "github.com/slack-go/slack"
)

func (s *Service) ProcessReceivedSlackMessage(_ context.Context, msg *slackapi.Msg) error {
	_, _, err := s.SlackClient.PostMessage(msg.Channel, slackapi.MsgOptionAsUser(false), slackapi.MsgOptionText(msg.Text, false) /*, slackapi.MsgOptionPostEphemeral(msg.User)*/)
	return err
}
