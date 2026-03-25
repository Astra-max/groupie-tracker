package tests

import (
	"groupie-tracker/models"
	"groupie-tracker/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

var mockArtists = []models.Artist{
	{ID: 1, Name: "Queen"},
}

var mockDates = []models.Dates{}
var mockLocations = []models.Locations{}



func TestHomeHandler_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist", nil)
	w := httptest.NewRecorder()

	handler := handlers.HomeHandler(mockDates, mockArtists, mockLocations)
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if w.Body.Len() == 0 {
		t.Errorf("expected response body, got empty")
	}
}

func TestArtistHandler_Valid(t *testing.T) {
	artists := []models.Artist{
		{ID: 1, Name: "Queen"},
	}

	relations := []models.Relation{
		{
			ID: 1,
			DatesLocations: map[string][]string{
				"dunedin-new_zealand": {"10-02-2020"},
			},
		},
	}

	dates := []models.Dates{}

	req := httptest.NewRequest(http.MethodGet, "/artist/1", nil)
	w := httptest.NewRecorder()

	handler := handlers.ArtistHandler(artists, relations, dates)
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestSearchResultsHandler(t *testing.T) {
	artists := []models.Artist{
		{ID: 1, Name: "Queen"},
		{ID: 2, Name: "SOJA"},
	}

	form := strings.NewReader("search=drake")
	req := httptest.NewRequest(http.MethodPost, "/search/results", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	handler := handlers.SearchResultsHandler(artists)
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}