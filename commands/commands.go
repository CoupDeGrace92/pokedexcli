package cmd

import (
	"os"
	"fmt"
	"github.com/CoupDeGrace92/pokedexcli/pokehttp"
	"github.com/CoupDeGrace92/pokedexcli/state"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*state.Config, ...string) error
}

func CommandExit(cfg *state.Config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *state.Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("SUPPORTED COMMANDS:")
		for str, cmd := range SupportedCommands {
			fmt.Printf("Command: %s\n", str)
			fmt.Printf("    name: %s\n", cmd.Name)
			fmt.Printf("    description: %s\n\n", cmd.Description)
		}
		return nil
	}
	for _, command := range args {
		cmd := SupportedCommands[command]
		fmt.Printf("Command: %s\n", command)
		fmt.Printf("    name: %s\n", cmd.Name)
		fmt.Printf("    description: %s\n\n", cmd.Description)
	}
	return nil
}

var SupportedCommands = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"map": {
		Name:        "map",
		Description: "Display the next 20 locations from the current location",
		Callback:    pokehttp.Map,
	},
	"mapb": {
		Name:        "map back",
		Description: "Display the previous 20 locations from the current location",
		Callback:    pokehttp.MapB,
	},
	"explore": {
		Name: "explore",
		Description: "Display the pokemon that can be encountered at the requested location(s)",
		Callback: pokehttp.Explore,
	},
}

//To avoid an initialization cycle, we only add this to supported commands after initializing the map.
func init() {
	SupportedCommands["help"] = CliCommand{
		Name:        "help",
		Description: "Lists the supported commands and thier functions if no args are listed, only lists args commands if args are listed",
		Callback:    CommandHelp,
	}

}
