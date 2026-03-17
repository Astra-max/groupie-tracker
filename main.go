package main

import (
    "fmt"
    "log"
    "net/http"
    "groupie-tracker/models"
    "groupie-tracker/handlers"
    "groupie-tracker/api"
)

func main() {
    fmt.Println("🎸 Groupie Tracker Server Starting...")
    fmt.Println("======================================")
    
    client := api.NewClient()
    artists, locations, dates, relations, err := client.GetAllData()
    
    if err != nil {
        fmt.Println("⚠️  Could not connect to API:", err)
        fmt.Println("📋 Using mock data instead...")
        artists = api.GetMockArtists()
        locations = []models.Locations{}
        dates = []models.Dates{}
        relations = []models.Relation{}
    }
    
    fmt.Printf("✅ Loaded %d artists\n", len(artists))
    fmt.Printf("✅ Loaded %d location sets\n", len(locations))
    fmt.Printf("✅ Loaded %d date sets\n", len(dates))
    fmt.Printf("✅ Loaded %d relation sets\n", len(relations))
    
    // Set up routes
    fmt.Println("\n🚀 Setting up routes...")
    http.HandleFunc("/", handlers.HomeHandler(artists))
    http.HandleFunc("/artist/", handlers.ArtistHandler(artists, relations))
    http.HandleFunc("/search", handlers.SearchPageHandler(artists))
    http.HandleFunc("/search/results", handlers.SearchResultsHandler(artists))
    
    port := ":8081"
    fmt.Printf("\n🌍 Server running on http://localhost%s\n", port)
    fmt.Println("📝 Press Ctrl+C to stop")
    fmt.Println("======================================")
    
    log.Fatal(http.ListenAndServe(port, nil))
}