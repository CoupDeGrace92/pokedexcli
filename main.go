package main

import (
	"os"
	"github.com/CoupDeGrace92/pokedexcli/repl"
	"github.com/CoupDeGrace92/pokedexcli/state"
	"fmt"
	"time"
)


func main() {

	cfg := &state.Config{
		Id:            0,
		LocationCache: nil,
		Interval:      45 * time.Second,
	}

	fmt.Print("Welcome to the Pokedex!\n")
	reader := os.Stdin
	repl.CommandReader(reader, cfg)
}