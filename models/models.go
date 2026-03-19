package models


// Events - get events with their time, date and location

type Event struct {
    Artist string
    Date string
    Location string
}

// Artist represents a band or artist's basic information
type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
}

// Locations stores all concert locations for an artist
type Locations struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
}

// Dates stores all concert dates for an artist
type Dates struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}

// Relation connects locations with dates
type Relation struct {
    ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

// Response wrappers for API (since locations, dates, relation come wrapped in an "index" field)
type LocationsResponse struct {
    Index []Locations `json:"index"`
}

type DatesResponse struct {
    Index []Dates `json:"index"`
}

type RelationResponse struct {
    Index []Relation `json:"index"`
}

// DateLocation represents a single concert event (one date at one location)
type DateLocation struct {
    Date     string
    Location string
}

// ArtistDetails combines all information for one artist
type ArtistDetails struct {
    Artist
    DatesLocations []DateLocation
}