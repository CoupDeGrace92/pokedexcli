package state
import (
	"github.com/CoupDeGrace92/pokedexcli/internal"
	"time"
)

type Config struct {
	Id             int
	LocationCache  *internal.Cache
	Interval       time.Duration
	PokeDex        map[string]Pokemon
	PokemonCache   *internal.Cache
}

type Pokemon struct {
	Name            string
	Url             string
	BaseExperience  int `json:"base_experience"`
	Height          int
	Weight          int
	Stats           []Stats
	Types           []Types
}

type Stats struct {
	BaseStat        int `json:"base_stat"`
	Stat            StatType            
}

type StatType struct {
	Name            string
}

type Types struct {
	Type            TypeType
}

type TypeType struct {
	Name            string
}