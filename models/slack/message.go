package slack

type MessageChannels struct {
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	ApiAppID    string   `json:"api_app_id"`
	Event       Event    `json:"event"`
	EventType   string   `json:"type"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedTeams []string `json:"authed_teams"`
}

type Event struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	User           string `json:"user"`
	Text           string `json:"text"`
	TimeStamp      string `json:"ts"`
	EventTimeStamp string `json:"event_ts"`
	ChannelType    string `json:"channel_type"`
}
