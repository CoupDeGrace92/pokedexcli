package state
import (
	"github.com/CoupDeGrace92/pokedexcli/internal"
	"time"
)

type Config struct {
	Id             int
	LocationCache  *internal.Cache
	Interval       time.Duration
}