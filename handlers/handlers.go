package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// PageData holds data to send to templates
type PageData struct {
	Title        string
	Artists      []models.Artist
	Artist       *models.ArtistDetails
	Search       string
	Dates        []models.Dates
	Results      []models.Artist
	SingleArtist *models.Artist
	Concert      *models.Event
	Error        string
	Mode         string
	AllConcerts []models.Event
	Locations []models.Locations

}

// HomeHandler displays the main page with all artists
func HomeHandler(dates []models.Dates, artists []models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var defaultArtist *models.ArtistDetails

		if len(artists) > 0 {
			defaultArtist = &models.ArtistDetails{
				Artist: artists[0],
			}
		}

		data := PageData{
			Title:   "Groupie Trackers - All Artists",
			Dates:   dates,
			Artists: artists,
			Artist:  defaultArtist,
			Mode:    "dash",
		}

		tmpl.Execute(w, data)
	}
}

// ArtistHandler displays details for a single artist
func ArtistHandler(artists []models.Artist, relations []models.Relation, dates []models.Dates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) != 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(pathParts[2])
		fmt.Println(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}

		var foundArtist *models.Artist
		for i := range artists {
			if artists[i].ID == id {
				foundArtist = &artists[i]
				break
			}
		}

		if foundArtist == nil {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		}

		artistDetails := &models.ArtistDetails{
			Artist:         *foundArtist,
			DatesLocations: []models.DateLocation{},
		}

		for _, rel := range relations {
			if rel.ID == id {
				for location, dates := range rel.DatesLocations {
					cleanLocation := strings.ReplaceAll(location, "-", ", ")
					cleanLocation = strings.Title(cleanLocation)

					for _, date := range dates {
						artistDetails.DatesLocations = append(artistDetails.DatesLocations, models.DateLocation{
							Date:     date,
							Location: cleanLocation,
						})
					}
				}
				break
			}
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:   foundArtist.Name + " - Groupie Trackers",
			Artist:  artistDetails,
			Artists: artists,
			Dates:   dates,
			Mode:    "home",
		}

		tmpl.Execute(w, data)
	}
}

// SearchPageHandler displays the search page
func SearchPageHandler(artists []models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/search.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:   "Search Artists - Groupie Trackers",
			Artists: artists,
		}

		tmpl.Execute(w, data)
	}
}

// SearchResultsHandler processes search and shows results
func SearchResultsHandler(artists []models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query := r.FormValue("search")
		query = strings.ToLower(strings.TrimSpace(query))

		var results []models.Artist

		if query != "" && len(query) >= 2 {
			for _, artist := range artists {
				if strings.Contains(strings.ToLower(artist.Name), query) {
					results = append(results, artist)
					continue
				}

				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), query) {
						results = append(results, artist)
						break
					}
				}
			}
		}

		tmpl, err := template.ParseFiles("templates/search.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:   "Search Results - Groupie Trackers",
			Artists: artists,
			Search:  query,
			Results: results,
		}

		tmpl.Execute(w, data)
	}
}
