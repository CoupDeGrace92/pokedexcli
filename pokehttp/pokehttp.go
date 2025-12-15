package pokehttp

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/CoupDeGrace92/pokedexcli/state"
	"github.com/CoupDeGrace92/pokedexcli/internal"
	"time"
	"io"
	"math/rand"
)

type LocationArea struct {
	Id                  int
	Name                string
	PokemonEncounters  []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon     state.Pokemon
	//Also version details but we are not implementing it now
}


func GetMapTest(id string) (LocationArea, error){
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", id)

	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("Error: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error: %v", err)
	}

	fmt.Println(string(body))
	
	var locationMap LocationArea
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationMap); err != nil{
		return LocationArea{}, fmt.Errorf("Error: %v", err)
	}

	return locationMap, nil
}

func GetMap(id string, cache *internal.Cache, interval time.Duration) (LocationArea, error){
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", id)

	if value, ok := cache.Get(url); ok{
		var locationArea LocationArea
		err := json.Unmarshal(value, &locationArea)
		if err != nil{
			return LocationArea{}, fmt.Errorf("Error unmarshalling data from cache: %v", err)
		}
		return locationArea, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error with get request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("Unexpected status code: %v", resp.StatusCode)
	}

	var locationMap LocationArea
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationMap); err != nil{
		return LocationArea{}, fmt.Errorf("Error with decoding json from pokeapi: %w", err)
	}

	//We need to convert locationMap to byte[]
	byteSlice, err:= json.Marshal(locationMap)
	if err != nil{
		return LocationArea{}, fmt.Errorf("Error with marshalling json from locationMap struct: %v", err)
	}
	cache.Add(url, byteSlice)

	return locationMap, nil
}

func Map(st *state.Config, args ...string) error {
	startId := st.Id
	if st.LocationCache == nil{
		c  := internal.NewCache(st.Interval)
		st.LocationCache = &c
	}
	cache := st.LocationCache
	interval := st.Interval
	for i:=startId+1; i<=startId+20; i++ {
		locId := strconv.Itoa(i)
		location, err := GetMap(locId, cache, interval)
		if err != nil{
			//fmt.Printf("%v", err)
			st.Id = i-1
			return nil
		}
		fmt.Println(location.Name)
	}
	st.Id = startId+20
	return nil
}

func MapB(st *state.Config, args ...string) error {
	startId := st.Id
	if st.LocationCache == nil{
		c  := internal.NewCache(st.Interval)
		st.LocationCache = &c
	}
	cache := st.LocationCache
	interval := st.Interval
	for i := startId-1; i>=startId-20; i-- {
		locId := strconv.Itoa(i)
		location, err := GetMap(locId, cache, interval)
		if err != nil{
			//fmt.Printf("%v", err)
			st.Id = i+1
			return nil
		}
		fmt.Println(location.Name)
	}
	st.Id = startId-20
	return nil
}

func Explore(st *state.Config, args ...string) error {
	if st.LocationCache == nil{
		c := internal.NewCache(st.Interval)
		st.LocationCache = &c
	}
	cache := st.LocationCache
	interval := st.Interval
	for _, loc := range args {
		location, err := GetMap(loc, cache, interval)
		if err != nil{
			fmt.Printf("Unable to find location %s: %v\n", loc, err)
			continue
		}
		fmt.Printf("Location: %s\n", loc)
		if len(location.PokemonEncounters) < 1{
			fmt.Printf("No pokemon found here\n")
			continue
		}
		for _, encounter := range location.PokemonEncounters {
			fmt.Printf("	%s\n", encounter.Pokemon.Name)
		}
	}
	return nil
}

func Catch(st *state.Config, args ...string) error {
	if len(args) == 0 {
		err := fmt.Errorf("No pokemon specified to catch\n")
		fmt.Println(err)
		return err
	}

	if st.PokemonCache == nil{
		c := internal.NewCache(st.Interval)
		st.PokemonCache = &c
	}

	for _, poke := range args {
		pokemon, err := GetPokemon(poke, st.PokemonCache, st.Interval)
		if err != nil{
			fmt.Printf("Pokemon %s not found: %v\n", poke, err)
			continue
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", poke)

		thresh := float64(pokemon.BaseExperience)/6.5
		rand := float64(rand.Intn(100)+1)
		if rand < thresh {
			fmt.Printf("%s escaped!\n", pokemon.Name)
			continue
		}

		fmt.Printf("%s was caught!\n", pokemon.Name)
		st.PokeDex[pokemon.Name] = pokemon
	}
	return nil
}

func GetPokemon(id string, cache *internal.Cache, interval time.Duration) (state.Pokemon, error){
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", id)

	if value, ok := cache.Get(url); ok{
		var poke state.Pokemon
		err := json.Unmarshal(value, &poke)
		if err != nil {
			return state.Pokemon{}, fmt.Errorf("Error unmarshalling data from cache: %v", err)
		}
		return poke, nil
	}

	resp, err := http.Get(url)
	if err != nil{
		return state.Pokemon{}, fmt.Errorf("Error with get request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return state.Pokemon{}, fmt.Errorf("Error with http status: %v", err)
	}

	var poke state.Pokemon
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&poke); err != nil{
		return state.Pokemon{}, fmt.Errorf("Error with decoding response: %v", err)
	}

	byteSlice, err := json.Marshal(poke)
	if err != nil{
		return state.Pokemon{}, fmt.Errorf("Error with marshalling json from poke: %v", err)
	}
	cache.Add(url, byteSlice)

	return poke, nil

}