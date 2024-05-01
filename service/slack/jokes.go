package slack

import (
	"fmt"
	"math/rand"

	"slackbot/api/openapi"

	"github.com/slack-go/slack"
)

type Joke struct {
	Setup     string
	Punchline string
}

var (
	jokes = []Joke{
		{
			Setup:     "What do you call a woman with one leg?.",
			Punchline: "Eileen",
		},
		{
			Setup:     "What did one hat say to the other?",
			Punchline: "You wait here. I’ll go on a head",
		},
		{
			Setup:     "What did the buffalo say when his son left for college?",
			Punchline: "Bison",
		},
		{
			Setup:     "What is an astronaut’s favorite part on a computer?",
			Punchline: "The space bar.",
		},
		{
			Setup:     "Why did the golfer wear two pairs of pants?",
			Punchline: "Just in case he got a hole in one!",
		},
	}
)

func (s *Service) pickAJoke() Joke {
	return jokes[rand.Int()%len(jokes)]
}

func (s *Service) tellAJoke(command openapi.CommandBody) error {
	joke := s.pickAJoke()

	_, _, err := s.client.PostMessage(command.ChannelID, slack.MsgOptionText(fmt.Sprintf("%s\n\n%s", joke.Setup, joke.Punchline), false), slack.MsgOptionPost())
	if err != nil {
		return err
	}

	return nil
}
