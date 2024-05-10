package service

import (
	"context"
	"fmt"

	"slackbot/api/model"

	"github.com/slack-go/slack"
)

type CommandWithDescription struct {
	Command     string
	Description string
}

var (
	commandsWithDescription = []CommandWithDescription{
		{Command: CommandGreet, Description: "Greets you in a very nice way"},
		{Command: CommandJoke, Description: "Tells you a pretty cheesy joke"},
		{Command: CommandPresentationStatus, Description: "Asks how a presentation is going"},
		{Command: CommandLogin, Description: "Trigger a login flow for Github"},
		{Command: CommandLogout, Description: "Logs out of Github"},
		{Command: CommandRepositories, Description: "Lists repositories of logged user"},
		{Command: CommandHelp, Description: "Shows this message"},
	}
)

func (s *Service) provideHelp(_ context.Context, command model.CommandBody) error {
	var blocks []slack.Block
	for _, comm := range commandsWithDescription {
		blocks = append(blocks, slack.SectionBlock{
			Type: "section",
			Fields: []*slack.TextBlockObject{
				{
					Type: "mrkdwn",
					Text: fmt.Sprintf("\"*%s*\" - %s", comm.Command, comm.Description),
				},
			},
		})
	}

	_, err := s.SlackClient.PostEphemeral(command.ChannelID, command.UserID, slack.MsgOptionBlocks(blocks...))
	return err
}
