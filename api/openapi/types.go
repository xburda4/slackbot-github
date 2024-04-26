package openapi

type RequestBodyMessage struct {
	Authorization string `header:"Authorization"`
	Body          MessageIM
}

type MessageIM struct {
	Token    string `json:"token" doc:"Token of the "`
	TeamId   string `json:"team_id" doc:"Team ID"`
	ApiAppId string `json:"api_app_id"`
	Event    struct {
		Type        string `json:"type"`
		Channel     string `json:"channel"`
		User        string `json:"user"`
		Text        string `json:"text" doc:"Text of the message" example:"This is a text"`
		Ts          string `json:"ts"`
		EventTs     string `json:"event_ts"`
		ChannelType string `json:"channel_type"`
	} `json:"event"`
	Type        string   `json:"type"`
	AuthedTeams []string `json:"authed_teams"`
	EventId     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
}
