package pokehttp

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type LocationArea struct {
	Id      int
	Name    string
}

func GetMap(id int) (locationMap LocationArea, err Error){
	idString := string(id)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", idString)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error with get request: %w", err)
	}
	defer res.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %v", resp.StatusCode)
	}

	var locationMap LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationMap); err != nil{
		return nil, fmt.Errorf("Error with decoding json from pokeapi: %w", err)
	}

	//Alternative to the decoder, we can unmarshall the data 
	//the distinction is decode streams the data into the json, unmarshalling creates a reader with all the data
	//and then unmarshalls all of it at once instead of as the byte data comes in
	//In small files, there is not much of a distinction, but decoding is better for large files or data that is in a constant stream
	//Marshalling is simpler for smaller files
	/*
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error with byte reader: %v", err)
	}
	var locationMap LocationArea
	if err := json.Unmarshal(bodyBytes, &locationMap); err = nil{
		return nil, fmt.Errorf("Error with unmarshalling json from pokeapi: %w", err)
	}
	*/


	return locationMap, nil
}

func Map(startId int) endId int {
	for i:=startID; i<startId+20; i++ {
		location, err := GetMap(i)
		if err != nil{
			endId := i
			//fmt.Printf("%v", err)
			return endId
		}
		fmt.Println(location.Name)
	}
	endId = i+19
	return endId
}

func MapB(startId int) endId int {
	for i := startID; i>startId-20; i-- {
		location, err := GetMap(i)
		if err != nil{
			endId :=i
			//fmt.Printf("%v", err)
			return endId
		}
		fmt.Println(location.Name)
	}
	endId = i-19
	return endId
}