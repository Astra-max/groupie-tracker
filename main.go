package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/api"
	"groupie-tracker/handlers"
	"groupie-tracker/models"
)

func main() {
	fmt.Println("Groupie Tracker Server Starting...")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	client := api.NewClient()

	artists, locations, dates, relations, err := client.GetAllData()

	if err != nil {
		fmt.Println("Could not connect to API:", err)
		fmt.Println("Using mock data instead...")
		artists = api.GetMockArtists()
		locations = []models.Locations{}
		dates = []models.Dates{}
		relations = []models.Relation{}
	}

	fmt.Printf("Loaded %d artists\n", len(artists))
	fmt.Printf("Loaded %d location sets\n", len(locations))
	fmt.Printf("Loaded %d date sets\n", len(dates))
	fmt.Printf("Loaded %d relation sets\n", len(relations))

	http.HandleFunc("/", handlers.HomeHandler(dates, artists, locations))
	http.HandleFunc("/artist/", handlers.ArtistHandler(artists, relations, dates))
	http.HandleFunc("/concerts/", handlers.ConcertsByDate(dates, artists, relations))
	http.HandleFunc("/search", handlers.SearchPageHandler(artists))
	http.HandleFunc("/search/results", handlers.SearchResultsHandler(artists))
	
	port := ":8000"
	fmt.Printf("\nServer running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
