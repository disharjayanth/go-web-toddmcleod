package main

import (
	"encoding/json"
	"fmt"
)

// Video struct
type Video struct {
	Width     int    `json:"Width"`
	Height    int    `json:"Height"`
	Title     string `json:"Title"`
	Thumbnail struct {
		URL    string `json:"Url"`
		Height int    `json:"Height"`
		Width  int    `json:"Width"`
	} `json:"Thumbnail"`
	Animated bool  `json:"Animated"`
	IDs      []int `json:"IDs"`
}

func main() {
	var video Video
	jsonString := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`
	err := json.Unmarshal([]byte(jsonString), &video)
	if err != nil {
		fmt.Println("Error Unmarshalling from JSON to Go code:", err)
		return
	}

	fmt.Println("Unmarshalled OP from JSON to GoCode: ", video)

	for key, value := range video.IDs {
		fmt.Println(key, ":", value)
	}

	fmt.Println("Nested struct:", video.Thumbnail.URL)
}
