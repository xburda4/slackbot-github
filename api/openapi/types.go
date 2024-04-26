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
