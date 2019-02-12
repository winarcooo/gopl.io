package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"humphrey bogart", "ingrid bergman"}},
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"humphrey bogart", "ingrid bergman"}},
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"humphrey bogart", "ingrid bergman"}},
	}

	// data , err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
