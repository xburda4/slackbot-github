package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

	"slackbot/api/openapi"

	"github.com/ajg/form"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/slack-go/slack"
)

func (h *Handler) SetupRoutes() *chi.Mux {

	mux := chi.NewMux()
	api := humachi.New(mux, huma.DefaultConfig("Slackbot", "0.1.0"))

	mux.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(os.Getenv("PUBLIC_FOLDER")))))

	huma.Register(api, huma.Operation{
		OperationID:   "receiveMessage",
		Method:        http.MethodPost,
		Path:          "/slack/events",
		Summary:       "Handles events from Slack",
		Description:   "Handles events from Slack",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
	}, h.postEventReq)

	huma.Register(api, huma.Operation{
		OperationID:   "interact",
		Method:        http.MethodPost,
		Path:          "/slack/interactive",
		Summary:       "Handles interactive events from Slack",
		Description:   "Handles interactive events from Slack",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
	}, h.postInteractiveReq)

	huma.Register(api, huma.Operation{
		OperationID:      "receiveCommand",
		Method:           http.MethodPost,
		SkipValidateBody: true,
		Path:             "/slack/command",
		Summary:          "Handles commands from Slack",
		Description:      "Handles commands from Slack",
		DefaultStatus:    http.StatusNoContent,
		Tags:             []string{"slack"},
	}, h.postCommandReq)

	huma.Register(api, huma.Operation{
		OperationID:   "finishAuthentication",
		Method:        http.MethodGet,
		Path:          "/slack/oauth",
		Summary:       "Finish authentication of the bot",
		Description:   "Finish authentication of the bot",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"slack"},
	}, h.handleFinishAuthRequest)

	huma.Register(api, huma.Operation{
		OperationID:   "finishGithubOauth",
		Method:        http.MethodGet,
		Path:          "/github/oauth",
		Summary:       "Finish GitHub authentication of the bot",
		Description:   "Finish GitHub authentication of the bot",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"github"},
	}, h.handleFinishGithubAuthRequest)

	return mux
}

// Get values from request body or path

/*func (h *Handler) handleAuthorizationReq(ctx huma.Context, _ *struct{}) (*struct{}, error) {
	redirectURL, err := h.service.SlackService.Authenticate()
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return nil, nil
}*/

func (h *Handler) handleFinishGithubAuthRequest(ctx context.Context, req *openapi.GithubOauthReq) (*openapi.GithubOauthResp, error) {
	err := h.service.GithubOauth(ctx, req.Code, req.State)
	if err != nil {
		return nil, err
	}

	return &openapi.GithubOauthResp{
		Status:      http.StatusOK,
		Body:        []byte("<html><head><script type=\"text/javascript\">window.close();</script></head><body><p>Operation successful. You can safely close this window.</p></body></html>"),
		ContentType: "text/html",
	}, nil
}

func (h *Handler) handleFinishAuthRequest(_ context.Context, _ *struct{}) (*struct{}, error) {
	//err := h.service.SlackService.Authorize()

	return nil, nil
}

func (h *Handler) postInteractiveReq(ctx context.Context, requestBody *openapi.InteractiveReq) (*openapi.InteractiveResp, error) {
	unescaped, err := url.QueryUnescape(string(requestBody.RawBody))
	if err != nil {
		return nil, err
	}

	unescapedJSON := strings.TrimPrefix(unescaped, "payload=")

	var ic slack.InteractionCallback
	if err := json.NewDecoder(strings.NewReader(unescapedJSON)).Decode(&ic); err != nil {
		return nil, err
	}

	if ic.Type != "block_actions" {
		return nil, nil
	}

	_, err = h.service.UpdateModalView(ctx, ic)
	if err != nil {
		return &openapi.InteractiveResp{
			Status: 200,
		}, nil
	}

	return &openapi.InteractiveResp{
		Status: 200,
	}, nil

}

func (h *Handler) postEventReq(ctx context.Context, requestBody *openapi.RequestBodyMessage) (*openapi.EventsResp, error) {
	if requestBody == nil {
		return nil, errors.New("request body is nil")
	}

	if requestBody.Body.Type == "url_verification" {
		challenge := requestBody.Body.Challenge

		return &openapi.EventsResp{
			Body: openapi.EventsRespBody{
				Challenge: challenge,
			},
			ContentType: "application/json",
		}, nil

	}

	if requestBody.Body.Event.BotID != "" {
		return nil, nil
	}
	msg := &slack.Msg{
		Channel: requestBody.Body.Event.Channel,
		User:    requestBody.Body.Event.User,
		Text:    requestBody.Body.Event.Text,
	}

	err := h.service.ProcessReceivedSlackMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &openapi.EventsResp{}, nil
}

func (h *Handler) postCommandReq(ctx context.Context, requestBody *openapi.CommandReq) (*struct{}, error) {
	if requestBody == nil {
		return nil, huma.Error400BadRequest("request body is nil")
	}

	var commandBody openapi.CommandBody
	// Get request info you don't normally have access to.
	d := form.NewDecoder(bytes.NewReader(requestBody.RawBody))

	if err := d.Decode(&commandBody); err != nil {
		return nil, huma.Error400BadRequest("invalid body")
	}

	err := h.service.HandleCommand(ctx, commandBody)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return nil, nil
}
