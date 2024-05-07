package openapi

import (
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
)

type RequestBodyMessage struct {
	TimeStamp      string `header:"X-Slack-Request-Timestamp"`
	SlackSignature string `header:"x-slack-signature"`
	Body           MessageIM
	RawBody        []byte
	_              struct{} `json:"-" additionalProperties:"true"`
}

func (rbm *RequestBodyMessage) Resolve(ctx huma.Context, prefix *huma.PathBuffer) []error {
	if err := json.NewDecoder(ctx.BodyReader()).Decode(&rbm.Body); err != nil {
		return nil
	}

	return nil
}

type MessageIM struct {
	Token string `json:"token,omitempty" doc:"Token of the "`
	Event struct {
		Type    string   `json:"type,omitempty" `
		Channel string   `json:"channel,omitempty"`
		User    string   `json:"user,omitempty"`
		Text    string   `json:"text,omitempty" doc:"Text of the message" example:"This is a text"`
		BotID   string   `json:"bot_id,omitempty"`
		_       struct{} `json:"-" additionalProperties:"true"`
	} `json:"event,omitempty"`
	Type      string   `json:"type"`
	Challenge string   `json:"challenge,omitempty" required:"false"`
	_         struct{} `json:"-" additionalProperties:"true"`
}

type InteractiveReq struct {
	//Body    slack.InteractionCallback
	RawBody []byte
}

type InteractiveResp struct {
	Status int
}

type EventsRespBody struct {
	Challenge string `json:"challenge"`
}

type EventsResp struct {
	Body        EventsRespBody
	ContentType string `header:"Content-Type"`
}

type CommandBody struct {
	Token               string `json:"token" form:"token"`
	TeamID              string `json:"team_id" form:"team_id"`
	TeamDomain          string `json:"team_domain" form:"team_domain"`
	ChannelID           string `json:"channel_id" form:"channel_id"`
	ChannelName         string `json:"channel_name" form:"channel_name"`
	UserID              string `json:"user_id" form:"user_id"`
	UserName            string `json:"user_name" form:"user_name"`
	Command             string `json:"command" form:"command"`
	Text                string `json:"text" form:"text"`
	ResponseURL         string `json:"response_url" form:"response_url"`
	TriggerID           string `json:"trigger_id" form:"trigger_id"`
	APIAppID            string `json:"api_app_id" form:"api_app_id"`
	EnterpriseID        string `json:"enterprise_id" form:"enterprise_id"`
	EnterpriseName      string `json:"enterprise_name" form:"enterprise_name"`
	IsEnterpriseInstall bool   `json:"is_enterprise_install" form:"is_enterprise_install"`
}

type CommandReq struct {
	TimeStamp      string `header:"X-Slack-Request-Timestamp"`
	SlackSignature string `header:"x-slack-signature"`
	ContentType    string `header:"Content-Type"`
	RawBody        []byte
}

type GithubOauthReq struct {
	Code  string `query:"code"`
	State string `query:"state"`
}

type GithubOauthResp struct {
	Status      int
	Body        []byte
	ContentType string `header:"Content-Type"`
}
