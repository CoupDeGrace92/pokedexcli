package main

import (
	"os"
	"github.com/CoupDeGrace92/pokedexcli/repl"
	"github.com/CoupDeGrace92/pokedexcli/state"
	"fmt"
)


func main() {
	cfg := &state.Config{
		Id: 0,
	}

	fmt.Print("Welcome to the Pokedex!\n")
	reader := os.Stdin
	repl.CommandReader(reader, cfg)
}