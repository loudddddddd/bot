package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"main/category"
)

// Basic command structure
type BotCommand struct {
	Name        string
	Description string
	Options     []discord.CommandOption
	Callback    func(ev *discord.InteractionEvent, _ *discord.CommandInteraction) *api.InteractionResponse
	Category    category.Category
}

// Command utilities
func ErrorResponse(err error) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content:         option.NewNullableString("**Error:** " + err.Error()),
			Flags:           discord.EphemeralMessage,
			AllowedMentions: &api.AllowedMentions{ /* none */ },
		},
	}
}

func Defer(flags discord.MessageFlags) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.DeferredMessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags: flags,
		},
	}
}
