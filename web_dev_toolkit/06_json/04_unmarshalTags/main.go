package main

import (
	"encoding/json"
	"fmt"
)

type cities []struct {
	Bali       string  `json:"Postal"`
	Kauai      float64 `json:"Latitude"`
	Maui       float64 `json:"Longitude"`
	Java       string  `json:"Address"`
	NewZealand string  `json:"City"`
	Skye       string  `json:"State"`
	Oahu       string  `json:"Zip"`
	Hawaii     string  `json:"Country"`
}

func main() {
	var places cities
	jsonCitiesString := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(jsonCitiesString), &places)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Array of Objects (JSON) to Slice of cities struct:", places)
}
