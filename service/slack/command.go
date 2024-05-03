package slack

import (
	"context"
	"strings"

	"slackbot/api/openapi"
)

const (
	CommandJoke         = "joke"
	CommandGreet        = "greet"
	CommandLogin        = "login"
	CommandLogout       = "logout"
	CommandRepositories = "repositories"
)

func (s *Service) HandleCommand(ctx context.Context, command openapi.CommandBody) error {
	commandText, _, isFound := strings.Cut(command.Text, " ")
	if !isFound {

	}

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

	default:
		//TODO: return 400
	}

	return nil
}
