package pokehttp

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/CoupDeGrace92/pokedexcli/state"
	"github.com/CoupDeGrace92/pokedexcli/internal"
	"time"
)

type LocationArea struct {
	Id      int
	Name    string
}

func GetMap(id int, cache *internal.Cache, interval time.Duration) (LocationArea, error){
	idString := strconv.Itoa(id)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", idString)

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

func Map(st *state.Config) error {
	startId := st.Id
	if st.LocationCache == nil{
		c  := internal.NewCache(st.Interval)
		st.LocationCache = &c
	}
	cache := st.LocationCache
	interval := st.Interval
	for i:=startId+1; i<=startId+20; i++ {
		location, err := GetMap(i, cache, interval)
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

func MapB(st *state.Config) error {
	startId := st.Id
	if st.LocationCache == nil{
		c  := internal.NewCache(st.Interval)
		st.LocationCache = &c
	}
	cache := st.LocationCache
	interval := st.Interval
	for i := startId-1; i>=startId-20; i-- {
		location, err := GetMap(i, cache, interval)
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