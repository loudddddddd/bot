package main

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	c "main/commands"
	"main/commands/misc"
	"os"
)

// To run, do `BOT_TOKEN="TOKEN HERE" go run .`

var commands = c.GetAllCommands()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("An error occured while loading godotenv")
	}
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	log.Println("Registering commands")
	c.AddCommand(misc.EchoCommand)
	log.Println("Registered commands!")

	// Set the commands var again so it gets new commands
	commands = c.GetAllCommands()

	var h handler
	h.s = state.New("Bot " + token)
	h.s.AddInteractionHandler(&h)
	h.s.AddIntents(gateway.IntentGuilds)
	h.s.AddHandler(func(*gateway.ReadyEvent) {
		me, _ := h.s.Me()
		log.Println("Connected to the gateway as", me.Tag())
	})

	if err := overwriteCommands(h.s); err != nil {
		log.Fatalln("Cannot update commands:", err)
	}

	if err := h.s.Connect(context.Background()); err != nil {
		log.Fatalln("Cannot connect:", err)
	}
}

func overwriteCommands(s *state.State) error {
	app, err := s.CurrentApplication()
	if err != nil {
		return errors.Wrap(err, "cannot get current app ID")
	}
	_, err = s.BulkOverwriteCommands(app.ID, commands)
	if err == nil {
		log.Println("Overwrote all commands!")
	}
	return err 
}

type handler struct {
	s *state.State
}

func (h *handler) HandleInteraction(ev *discord.InteractionEvent) *api.InteractionResponse {
	switch data := ev.Data.(type) {
	case *discord.CommandInteraction:
		switch data.Name {
		default:
			//return ErrorResponse(fmt.Errorf("unknown command %q", data.Name))
			for _, cmd := range c.GetAllCommandsRaw() {
				if cmd.Name == data.Name {
					return cmd.Callback(ev, data)
				}
			}
			return ErrorResponse(fmt.Errorf("unknown command %q", data.Name))
		}
	default:
		return ErrorResponse(fmt.Errorf("unknown interaction %T", ev.Data))
	}
}


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

