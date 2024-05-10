package service

import (
	"context"
	"fmt"
	"strings"

	"slackbot/api/model"

	"github.com/slack-go/slack"
)

const (
	CommandJoke               = "joke"
	CommandGreet              = "greet"
	CommandLogin              = "login"
	CommandLogout             = "logout"
	CommandRepositories       = "repos"
	CommandPresentationStatus = "status"
	CommandHelp               = "help"
)

func (s *Service) HandleCommand(ctx context.Context, command model.CommandBody) error {
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
	case CommandHelp:
		err := s.provideHelp(ctx, command)
		if err != nil {
			return err
		}
	default:
		_, _, err := s.SlackClient.PostMessage(command.ChannelID,
			slack.MsgOptionText(fmt.Sprintf("The command you entered is unknown. You can try `help` command to check which commands are available."), false))
		if err != nil {
			return err
		}
	}

	return nil
}
