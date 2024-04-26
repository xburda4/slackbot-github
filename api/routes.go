package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"slackbot/api/openapi"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/slack-go/slack"
)

func (h *Handler) SetupRoutes() *chi.Mux {

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
		Path:          "/slack/events",
		Summary:       "Handles events from Slack",
		Description:   "Handles events from Slack",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
	}, h.postCommandReq)

	return mux
}

// Get values from request body or path

func (h *Handler) postCommandReq(ctx context.Context, requestBody *openapi.RequestBodyMessage) (*openapi.EventsResp, error) {
	if requestBody == nil {
		return nil, errors.New("request body is nil")
	}
	fmt.Println("here")

	if requestBody.Body.Type == "url_verification" {
		challenge := requestBody.Body.Challenge

		return &openapi.EventsResp{
			Body: openapi.EventsRespBody{
				Challenge: challenge,
			},
			ContentType: "application/json",
		}, nil

	}

	msg := &slack.Msg{
		Channel: requestBody.Body.Event.Channel,
		User:    requestBody.Body.Event.User,
		Text:    requestBody.Body.Event.Text,
	}

	err := h.service.SlackService.ProcessReceivedSlackMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &openapi.EventsResp{}, nil
}
