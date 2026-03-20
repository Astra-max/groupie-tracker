package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io/ioutil"
	"net/http"
	"time"
)

// Client handles all communication with the API
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		BaseURL: "https://groupietrackers.herokuapp.com/api",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// FetchArtists gets all artists from the API
func (c *Client) FetchArtists() ([]models.Artist, error) {
	var artists []models.Artist
	url := fmt.Sprintf("%s/artists", c.BaseURL)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return artists, nil
}

// FetchLocations gets all locations from the API
func (c *Client) FetchLocations() ([]models.Locations, error) {
	var response models.LocationsResponse
	url := fmt.Sprintf("%s/locations", c.BaseURL)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return response.Index, nil
}

// FetchDates gets all dates from the API
func (c *Client) FetchDates() ([]models.Dates, error) {
	var response models.DatesResponse
	url := fmt.Sprintf("%s/dates", c.BaseURL)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dates: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return response.Index, nil
}

// FetchRelations gets all relations from the API
func (c *Client) FetchRelations() ([]models.Relation, error) {
	var response models.RelationResponse
	url := fmt.Sprintf("%s/relation", c.BaseURL)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relations: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return response.Index, nil
}

// GetAllData fetches everything from the API in one call
func (c *Client) GetAllData() ([]models.Artist, []models.Locations, []models.Dates, []models.Relation, error) {
	artists, err := c.FetchArtists()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	locations, err := c.FetchLocations()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	dates, err := c.FetchDates()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	relations, err := c.FetchRelations()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return artists, locations, dates, relations, nil
}

// GetMockArtists returns sample artist data when API is unavailable
func GetMockArtists() []models.Artist {
	return []models.Artist{
		{
			ID:           1,
			Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
			Name:         "Queen",
			Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
			CreationDate: 1970,
			FirstAlbum:   "14-12-1973",
		},
		{
			ID:           2,
			Image:        "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
			Name:         "SOJA",
			Members:      []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"},
			CreationDate: 1997,
			FirstAlbum:   "05-06-2002",
		},
	}
}
