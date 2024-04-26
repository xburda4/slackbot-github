package api

import (
	"context"
	"net/http"

	"slackbot/api/openapi"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/slack-go/slack"
)

func SetupRoutes() *chi.Mux {

	mux := chi.NewMux()
	api := humachi.New(mux, huma.DefaultConfig("Slackbot", "0.1.0"))

	/*huma.Register(api, huma.Operation{
		OperationID:   "receiveCommand",
		Method:        http.MethodPost,
		Path:          "/slack/command",
		Summary:       "Handles slash command from Slack",
		Description:   "Handles slash command from Slack",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
		RequestBody: &huma.RequestBody{
			Description: "Slack command ",
			Content: map[string]*huma.MediaType{
				"application/json": {
					Schema:     nil,
					Example:    nil,
					Examples:   nil,
					Encoding:   nil,
					Extensions: nil,
				},
			},
			Required: true,
		},
	}, postCommandReq)*/

	huma.Register(api, huma.Operation{
		OperationID:   "receiveMessage",
		Method:        http.MethodPost,
		Path:          "/slack/message",
		Summary:       "Handles messages from Slack",
		Description:   "Handles messages from Slack",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
	}, postCommandReq)

	return mux
}

// Get values from request body or path

func postCommandReq(ctx context.Context, cmd *openapi.RequestBodyMessage) (*slack.SlashCommand, error) {
	return nil, nil
}
