package handlers

import (
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strings"
	"strconv"
)


func ConcertsByDate(dates []models.Dates,artists []models.Artist) http.HandlerFunc {
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

		arstistId := strings.Split(r.URL.Query().Get("date"), "|")

		userId, errId := strconv.Atoi(arstistId[0])

		if errId != nil {
			http.Error(w, "Invalid artistId", http.StatusBadRequest)
			return
		}

		var singleArtist models.Artist

		for _, art := range artists {
			if art.ID == userId {
				singleArtist = art
			}
		}

		defaultArtist = &models.ArtistDetails{
				Artist: singleArtist,
			}

		data := PageData{
			Title:   "Groupie Trackers - All Artists",
			Dates: dates,
			Artists: artists,
			Artist:  defaultArtist,
			Mode: "concert",
		}

		tmpl.Execute(w, data)
	}
}

func PreviousEvents(w http.ResponseWriter, req *http.Request) {}

func AllEvents(w http.ResponseWriter, req *http.Request) {}
