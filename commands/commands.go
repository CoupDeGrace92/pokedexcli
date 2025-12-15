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

func CommandInspect(cfg *state.Config, args ...string) error{
	if len(args) == 0 {
		fmt.Printf("No pokemon were entered\n")
		return nil
	}
	for _, poke := range args{
		pokeObject, ok := cfg.PokeDex[poke]
		if !ok {
			fmt.Printf("You have not caught that pokemon\n")
			continue
		}
		fmt.Printf("\nHeight: %v\n", pokeObject.Height)
		fmt.Printf("Weight: %v\n", pokeObject.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokeObject.Stats {
			fmt.Printf("	-%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, typ := range pokeObject.Types {
			fmt.Printf("	-%s\n", typ.Type.Name)
		}
	}
	fmt.Printf("\n")
	return nil
}

func CommandPokedex(cfg *state.Config, args ...string) error{
	fmt.Printf("Your Pokedex:\n")
	if len(cfg.PokeDex)==0{
		fmt.Printf("No Pokemon caught, get on catching ya ignoramus\n")
	}
	for name, _ := range cfg.PokeDex {
		fmt.Printf("	-%s\n",name)
	}
	fmt.Println()
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
		Name:        "explore",
		Description: "Display the pokemon that can be encountered at the requested location(s)",
		Callback:    pokehttp.Explore,
	},
	"catch": {
		Name:        "catch",
		Description: "Attempts to catch the specified pokemon, if successful, add it to the PokeDex",
		Callback:    pokehttp.Catch,
	},
	"inspect": {
		Name:         "inspect",
		Description:  "Displays a pokemons information if previously caught",
		Callback:     CommandInspect,
	},
	"pokedex": {
		Name:         "pokedex",
		Description:  "Displays the name of each pokemon caught",
		Callback:     CommandPokedex,
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
