package misc

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"main/category"
	"main/commands"
)

var EchoCommand commands.BotCommand = commands.BotCommand{
	Name:        "echo",
	Description: "echo back the argument",
	Options: []discord.CommandOption{
		&discord.StringOption{
			OptionName:  "argument",
			Description: "what's echoed back",
			Required:    true,
		},
	},
	Callback: func(ev *discord.InteractionEvent, data *discord.CommandInteraction) *api.InteractionResponse {
		var options struct {
			Arg string `discord:"argument"`
		}

		if err := data.Options.Unmarshal(&options); err != nil {
			return commands.ErrorResponse(err)
		}

		return &api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Content:         option.NewNullableString(options.Arg),
				AllowedMentions: &api.AllowedMentions{},
			},
		}
	},

	Category: category.MISC,
}
