package service

import (
	"context"

	"github.com/slack-go/slack"
)

type Servicer interface {
	SlackServicer
}

type SlackServicer interface {
	ProcessReceivedSlackMessage(ctx context.Context, msg *slack.Msg) error
}
