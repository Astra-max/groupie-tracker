package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strings"
)

func ConcertsByDate(
	dates []models.Dates,
	artists []models.Artist,
	relations []models.Relation,
) http.HandlerFunc {

	// Preprocess relations into a map for O(1) lookup
	relMap := make(map[int]models.Relation, len(relations))
	for _, rel := range relations {
		relMap[rel.ID] = rel
	}

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

		rawDate := r.URL.Query().Get("date")
		if rawDate == "" {
			http.Error(w, "Missing date parameter", http.StatusBadRequest)
			return
		}
		cleanDate := strings.TrimPrefix(rawDate, "*")

		var concerts []models.Event

		// Loop over artists once, lookup their relations
		for _, artist := range artists {
			if rel, ok := relMap[artist.ID]; ok {
				var matchingLocations []string
				for loc, dateList := range rel.DatesLocations {
					for _, d := range dateList {
						if strings.TrimPrefix(d, "*") == cleanDate {
							matchingLocations = append(matchingLocations, loc)
						}
					}
				}
				if len(matchingLocations) > 0 {
					concerts = append(concerts, models.Event{
						Name:     artist.Name,
						Image:    artist.Image,
						Date:     rawDate,
						Location: matchingLocations,
					})
				}
			}
		}

		if len(concerts) == 0 {
			http.Error(w, "No concerts found for this date", http.StatusNotFound)
			return
		}

		data := PageData{
			Title:       fmt.Sprintf("Concerts on %s", rawDate),
			Dates:       dates,
			Artists:     artists,
			Concert:     &concerts[0],
			Mode:        "concerts-list",
			AllConcerts: concerts,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
			return
		}
	}
}

func PreviousEvents(w http.ResponseWriter, req *http.Request) {}

func AllEvents(w http.ResponseWriter, req *http.Request) {}
