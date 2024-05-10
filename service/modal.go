package service

import (
	"context"

	"slackbot/api/model"

	"github.com/slack-go/slack"
)

const (
	ModalStatusPrettyWell = "pretty-well"
	ModalStatusSob        = "sob"
	ModalStatusAllGood    = "all-good"
	ModalStatusNeutral    = "neutral"
	ModalStatusCrying     = "crying"
)

func (s *Service) askForPresentationStatus(_ context.Context, command model.CommandBody) error {
	_, err := s.SlackClient.OpenView(command.TriggerID, slack.ModalViewRequest{
		Type:       "modal",
		CallbackID: "presentation-status",
		Title:      slack.NewTextBlockObject("plain_text", "Mental health check", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.SectionBlock{
					Type: "section",
					Text: &slack.TextBlockObject{
						Type: "plain_text",
						Text: "How is the presentation going?",
					},
					BlockID: "section-status",
				},
				slack.ActionBlock{
					Type: "actions",
					Elements: &slack.BlockElements{
						ElementSet: []slack.BlockElement{
							slack.ButtonBlockElement{
								Type: "button",
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: ":100:",
								},
								ActionID: ModalStatusPrettyWell,
							},
							slack.ButtonBlockElement{
								Type: "button",
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: ":white_check_mark:",
								},
								ActionID: ModalStatusAllGood,
							},
							slack.ButtonBlockElement{
								Type: "button",
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: ":neutral_face:",
								},
								ActionID: ModalStatusNeutral,
							},
							slack.ButtonBlockElement{
								Type: "button",
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: ":cry:",
								},
								ActionID: ModalStatusCrying,
							},
							slack.ButtonBlockElement{
								Type: "button",
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: ":sob:",
								},
								ActionID: ModalStatusSob,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateModalView(_ context.Context, body slack.InteractionCallback) (shouldClose bool, err error) {
	switch body.View.CallbackID {
	case "presentation-status":
		if len(body.ActionCallback.BlockActions) == 0 {
			shouldClose = true
			return
		}

		switch body.ActionCallback.BlockActions[0].ActionID {
		case ModalStatusSob:
			_, err := s.SlackClient.UpdateView(slack.ModalViewRequest{
				Type:       "modal",
				CallbackID: "presentation-status",
				Title:      slack.NewTextBlockObject("plain_text", "Mental health check", false, false),
				Blocks: slack.Blocks{
					BlockSet: []slack.Block{
						slack.SectionBlock{
							Type: "section",
							Text: &slack.TextBlockObject{
								Type: "plain_text",
								Text: "Just point at the person making you sad!",
							},
						},
					},
				},
			}, body.View.ExternalID, body.Hash, body.View.ID)
			if err != nil {
				return false, err
			}
			if s.Socket != nil {
				(*s.Socket).Write([]byte("sob-reaction"))
			}
		case ModalStatusPrettyWell:
			_, err := s.SlackClient.UpdateView(slack.ModalViewRequest{
				Type:       "modal",
				CallbackID: "presentation-status",
				Title:      slack.NewTextBlockObject("plain_text", "Mental health check", false, false),
				Blocks: slack.Blocks{
					BlockSet: []slack.Block{
						slack.SectionBlock{
							Type: "section",
							Text: &slack.TextBlockObject{
								Type: "plain_text",
								Text: "Glad everything's working all right. You can mention a person making you feel safe.",
							},
						},
					},
				},
			}, body.View.ExternalID, body.Hash, body.View.ID)
			if err != nil {
				return false, err
			}
			if s.Socket != nil {
				(*s.Socket).Write([]byte("pretty-well-reaction"))
			}
		case ModalStatusNeutral:
			_, err := s.SlackClient.UpdateView(slack.ModalViewRequest{
				Type:       "modal",
				CallbackID: "presentation-status",
				Title:      slack.NewTextBlockObject("plain_text", "Mental health check", false, false),
				Blocks: slack.Blocks{
					BlockSet: []slack.Block{
						slack.SectionBlock{
							Type: "section",
							Text: &slack.TextBlockObject{
								Type: "plain_text",
								Text: "It will be better in the next few minutes.",
							},
						},
					},
				},
			}, body.View.ExternalID, body.Hash, body.View.ID)
			if err != nil {
				return false, err
			}
			if s.Socket != nil {
				(*s.Socket).Write([]byte("neutral-reaction"))
			}
		default:
			_, err := s.SlackClient.UpdateView(slack.ModalViewRequest{
				Type:         "modal",
				CallbackID:   "presentation-status",
				Title:        slack.NewTextBlockObject("plain_text", "Mental health check", false, false),
				ClearOnClose: true,
			}, body.View.ExternalID, body.Hash, body.View.ID)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}

	return
}
