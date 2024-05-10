package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"slackbot/api/model"

	"entgo.io/ent/dialect/sql"
	"github.com/slack-go/slack"
)

const (
	reposURL = "https://api.github.com/user/repos"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}

func (s *Service) listGithubRepos(ctx context.Context, command model.CommandBody) error {
	ghUser, err := s.Database.GithubUser.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.EQ("slack_id", command.UserID))
		}).Only(ctx)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, reposURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ghUser.GhAccessToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Accept", "application/vnd.github+json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var repositories []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		return err
	}

	repos := make([]slack.Block, 0, len(repositories))
	for _, repo := range repositories {
		repos = append(repos, slack.SectionBlock{
			Type: "section",
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*%s* - %s\nThis repository is %s", repo.Name, repo.Description, repo.Visibility),
			},
		})
	}
	_, err = s.SlackClient.PostEphemeral(command.ChannelID, command.UserID, slack.MsgOptionBlocks(repos...))
	if err != nil {
		return err
	}

	return nil
}
