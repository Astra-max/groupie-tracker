package handlers

import (
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strings"
	"strconv"
)


func ConcertsByDate(
	dates []models.Dates,
	artists []models.Artist,
	relations []models.Relation,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		raw := r.URL.Query().Get("date")
		parts := strings.Split(raw, "|")

		if len(parts) < 2 {
			http.Error(w, "Missing artistId or date", http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(parts[0])
		if err != nil {
			http.Error(w, "Invalid artistId", http.StatusBadRequest)
			return
		}

		selectedDate := parts[1]

		var singleArtist models.Artist
		found := false

		for _, art := range artists {
			if art.ID == userId {
				singleArtist = art
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		}

		var artistRelation models.Relation

		for _, rel := range relations {
			if rel.ID == userId {
				artistRelation = rel
				break
			}
		}

		var locations []string

		for date, locs := range artistRelation.DatesLocations {

			cleanDate := strings.TrimPrefix(date, "*")

			if cleanDate == selectedDate {
				locations = locs
				break
			}
		}

		if len(locations) == 0 {
			http.Error(w, "No locations found", http.StatusNotFound)
			return
		}

		concert := &models.Event{
			Name:     singleArtist.Name,
			Image:    singleArtist.Image,
			Date:     selectedDate,
			Location: strings.Join(locations, ", "),
		}

		data := PageData{
			Title:     "Groupie Trackers - Concerts",
			Dates:     dates,
			Artists:   artists,
			Concert:   concert,
			Mode:      "concert",
		}


		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
			return
		}
	}
}

func PreviousEvents(w http.ResponseWriter, req *http.Request) {}

func AllEvents(w http.ResponseWriter, req *http.Request) {}
