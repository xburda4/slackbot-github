package openapi

type RequestBodyMessage struct {
	TimeStamp      string    `header:"X-Slack-Request-Timestamp"`
	SlackSignature string    `header:"x-slack-signature"`
	Body           MessageIM `required:"false"`
	RawBody        []byte
}

type MessageIM struct {
	Token string `json:"token" doc:"Token of the "`
	Event struct {
		Type    string `json:"type" `
		Channel string `json:"channel"`
		User    string `json:"user"`
		Text    string `json:"text" doc:"Text of the message" example:"This is a text" required:"false"`
	} `json:"event,omitempty"`
	Type      string `json:"type"`
	Challenge string `json:"challenge,omitempty"`
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
