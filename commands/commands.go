package cmd

import (
	"os"
	"fmt"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func CommandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func CommandHelp(cmds map[string]CliCommand) error {
	fmt.Println("SUPPORTED COMMANDS:")
	for str, cmd := range cmds {
		fmt.Printf("Command: %s\n", str)
		fmt.Printf("    name: %s\n", cmd.Name)
		fmt.Printf("    description: %s\n", cmd.Description)
	}
	return nil
}

var SupportedCommands = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
}

//To avoid an initialization cycle, we only add this to supported commands after initializing the map.
func init() {
	SupportedCommands["help"] = CliCommand{
		Name:        "help",
		Description: "Lists the supported commands and thier functions",
		Callback:    func() error {
			return CommandHelp(SupportedCommands)
		},
	}
}
