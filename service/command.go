package service

import (
	"context"
	"fmt"
	"strings"

	"slackbot/api/openapi"

	"github.com/slack-go/slack"
)

const (
	CommandJoke               = "joke"
	CommandGreet              = "greet"
	CommandLogin              = "login"
	CommandLogout             = "logout"
	CommandRepositories       = "repositories"
	CommandPresentationStatus = "status"
)

func (s *Service) HandleCommand(ctx context.Context, command openapi.CommandBody) error {
	commandText, _, _ := strings.Cut(command.Text, " ")

	switch commandText {
	case CommandJoke:
		err := s.tellAJoke(command)
		if err != nil {
			return err
		}
	case CommandGreet:
		err := s.greet(command)
		if err != nil {
			return err
		}
	case CommandLogin:
		err := s.githubLogin(ctx, command)
		if err != nil {
			return err
		}
	case CommandLogout:
		err := s.githubLogout(ctx, command)
		if err != nil {
			return err
		}
	case CommandRepositories:
		err := s.listGithubRepos(ctx, command)
		if err != nil {
			return err
		}
	case CommandPresentationStatus:
		err := s.askForPresentationStatus(ctx, command)
		if err != nil {
			return err
		}
	default:
		_, _, err := s.SlackClient.PostMessage(command.ChannelID,
			slack.MsgOptionText(fmt.Sprintf("The command you entered is unknown"), false),
			slack.MsgOptionPost())
		if err != nil {
			return err
		}
	}

	return nil
}
