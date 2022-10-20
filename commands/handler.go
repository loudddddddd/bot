package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"log"
)

var Commands []BotCommand
var CommandData []api.CreateCommandData

func AddCommand(command BotCommand) {
	Commands = append(Commands, command)
	var comman2 api.CreateCommandData = api.CreateCommandData{
		Name:        command.Name,
		Description: command.Description,
		Options:     command.Options,
	}
	CommandData = append(CommandData, comman2)
	log.Println("Registered " + command.Name)
}

func GetAllCommands() []api.CreateCommandData {
	return CommandData
}

func GetAllCommandsRaw() []BotCommand {
	return Commands
}
