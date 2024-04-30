package slack

import (
	"strings"

	"slackbot/api/openapi"
)

const (
	CommandJoke         = "joke"
	CommandGreet        = "greet"
	CommandSing         = "sing"
	CommandLogin        = "login"
	CommandLogout       = "logout"
	CommandRepositories = "repositories"
)

func (s *Service) HandleCommand(command openapi.CommandBody) error {
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
		//TODO: call github
	case CommandSing:

	case CommandLogin:
	case CommandLogout:
	case CommandRepositories:
	default:
		//TODO: return 400
	}

	return nil
}
