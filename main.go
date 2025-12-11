package main

import (
	"os"
	"github.com/CoupDeGrace92/pokedexcli/repl"
	"fmt"
)

func main() {
	fmt.Print("Welcome to the Pokedex!\n")
	reader := os.Stdin
	repl.CommandReader(reader)
}