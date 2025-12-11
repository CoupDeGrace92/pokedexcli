package pokehttp

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/CoupDeGrace92/pokedexcli/state"
)

type LocationArea struct {
	Id      int
	Name    string
}

func GetMap(id int) (LocationArea, error){
	idString := strconv.Itoa(id)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", idString)

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

	//Alternative to the decoder, we can unmarshall the data 
	//the distinction is decode streams the data into the json, unmarshalling creates a reader with all the data
	//and then unmarshalls all of it at once instead of as the byte data comes in
	//In small files, there is not much of a distinction, but decoding is better for large files or data that is in a constant stream
	//Marshalling is simpler for smaller files
	/*
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error with byte reader: %v", err)
	}
	var locationMap LocationArea
	if err := json.Unmarshal(bodyBytes, &locationMap); err = nil{
		return LocationArea{}, fmt.Errorf("Error with unmarshalling json from pokeapi: %w", err)
	}
	*/


	return locationMap, nil
}

func Map(st *state.Config) error {
	startId := st.Id
	for i:=startId+1; i<=startId+20; i++ {
		location, err := GetMap(i)
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
	for i := startId-1; i>=startId-20; i-- {
		location, err := GetMap(i)
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