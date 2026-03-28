
# groupie-tracker

## Summary
The `groupie-tracker` project is a Go-based web application designed to fetch, process, and display comprehensive information about musical artists, their concert locations, and tour dates from an external API. It provides a user-friendly interface that allows users to browse artists, view detailed artist profiles, search for specific artists by name or member, and filter concert events by date. Built with Go's standard library for web serving and HTML templates for dynamic content, the application offers a clean and responsive user experience.

## Key Features
*   **External API Integration**: Seamlessly fetches artist, location, date, and relation data from the Groupie Tracker API.
*   **Concurrent Data Fetching**: Efficiently retrieves all necessary data using concurrent API calls to improve performance.
*   **Artist Management**: Displays a comprehensive list of all artists and provides detailed individual artist pages with concert information, formation dates, and members.
*   **Search Functionality**: Enables users to search for artists by their name or the names of their band members.
*   **Event Filtering**: Allows users to filter and display concert events based on specific dates.
*   **Dynamic UI**: Renders dynamic content through Go HTML templates, ensuring an interactive and rich user experience.
*   **Static Asset Serving**: Serves CSS stylesheets and SVG icons to enhance the application's visual appeal.
*   **Robust Error Handling**: Provides custom error pages for unhandled server errors and invalid requests, ensuring graceful degradation.
*   **DDL Utility (Internal)**: Includes internal utility functions for processing DDL statements, generating mock data, and replacing markers for SQL script generation, likely for development or testing purposes.

## Tech stack
-   **Primary Language**: Go
-   **Web Framework**: Go's `net/http` package
-   **Templating**: Go HTML/Template
-   **Data Interchange**: JSON
-   **Styling**: CSS, Material Symbols Outlined (Google Fonts)
-   **Utilities**: Python (for DDL processing, as indicated by docstrings)

## Installation
1.  **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/groupie-tracker.git
    cd groupie-tracker
    ```
2.  **Get Go dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Run the application**:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8000`.

## Folder Structure
```
groupie-tracker/
├── .gitignore               # Specifies intentionally untracked files to ignore
├── api/                     # Contains the API client for external data fetching
│   └── client.go            # Implementation of the Groupie Tracker API client
├── go.mod                   # Go module definition file for dependencies
├── handlers/                # HTTP request handlers for various routes
│   ├── handlers.go          # Core handlers for home page and artist details
│   ├── handlers2.go         # Additional handlers like search and date-based filtering
│   └── response.go          # Utility functions for HTTP responses, e.g., JSON decoding
├── logs/                    # Directory for server logging utilities
│   └── server-logs.go       # Placeholder for server logging functions
├── main.go                  # Main entry point of the application, initializes server and routes
├── models/                  # Defines data structures (structs) for the application
│   └── models.go            # Go structs for Artist, Location, Date, Relation, etc.
├── server/                  # Server configuration and startup logic
│   ├── server.go            # Server initialization and run methods
│   └── types.go             # Defines server-specific types (e.g., ServerConfig)
├── static/                  # Contains static assets served directly by the web server
│   ├── css/                 # Stylesheets for the web application
│   │   ├── artist.css       # Styles specific to artist detail pages
│   │   ├── dynamic-ui.css   # Styles for dynamic UI elements
│   │   ├── footer.css       # Styles for the website footer
│   │   ├── index.css        # General styles for the main pages
│   │   └── search.css       # Styles for the search page
│   └── icons/               # SVG icons used throughout the application
│       ├── cal.svg
│       ├── home.svg
│       ├── location.svg
│       └── m-icon.svg
├── templates/               # HTML templates rendered by the server
│   ├── index.html           # Main template for home page and artist details
│   ├── search.html          # Template for the artist search page
│   └── server-error.html    # Template for displaying server error messages
├── tests/                   # Contains unit and integration tests
│   ├── api_test.go          # Tests for the API client functions
│   ├── methods_test.go      # Tests for utility methods and handlers
│   └── templates/           # Duplicate templates used specifically for testing purposes
│       ├── index.html
│       ├── search.html
│       └── server-error.html
└── util/                    # General utility functions
    └── urls.go              # Functions related to URL handling
```

## API Documentation

The `groupie-tracker` application exposes several HTTP endpoints and utilizes an internal API client to interact with an external Groupie Tracker service.

### HTTP Handlers

These are the primary routes handled by the `groupie-tracker` web server:

*   **`HomeHandler`**
    *   **Description**: Displays the main page with a grid of all artists.
    *   **Path**: `/`
    *   **Method**: `GET`

*   **`ArtistHandler`**
    *   **Description**: Displays detailed information for a single artist, including concert dates and locations.
    *   **Path**: `/artist/{id}`
    *   **Method**: `GET`
    *   **Path Parameters**: `id` (integer, the artist's unique identifier)

*   **`SearchPageHandler`**
    *   **Description**: Renders the dedicated search interface for artists.
    *   **Path**: `/search`
    *   **Method**: `GET`

*   **`SearchResultsHandler`**
    *   **Description**: Processes search queries submitted via a `POST` request and displays matching artists.
    *   **Path**: `/search`
    *   **Method**: `POST`
    *   **Request Body**: Expects form data with search criteria (e.g., artist name, member name).

*   **`ConcertsByDate`**
    *   **Description**: Filters and displays concert events by a specified date.
    *   **Path**: `/concerts`
    *   **Method**: `GET`
    *   **Query Parameters**: `date` (string, e.g., `YYYY-MM-DD`)

*   **`ServerError`**
    *   **Description**: A generic handler to display a custom error page for unhandled server errors.

### Internal API Client (`api/client.go`)

This client is responsible for interacting with the external Groupie Tracker API.

*   **`type Client struct`**
    *   **Description**: Manages all communication with the external Groupie Tracker API, providing methods to fetch various data related to artists, locations, dates, and relations.

*   **`func NewClient() *Client`**
    *   **Description**: Initializes a new API client instance with the `BaseURL` and configures an `HTTP` client with a timeout.

*   **`func (c *Client) FetchArtists() ([]models.Artist, error)`**
    *   **Description**: Retrieves all artist data from the external API.

*   **`func (c *Client) FetchLocations() ([]models.Locations, error)`**
    *   **Description**: Retrieves all location data from the external API.

*   **`func (c *Client) FetchDates() ([]models.Dates, error)`**
    *   **Description**: Retrieves all date data from the external API.

*   **`func (c *Client) FetchRelations() ([]models.Relation, error)`**
    *   **Description**: Retrieves all relation data (linking artists to dates and locations) from the external API.

*   **`func (c *Client) GetAllData() ([]models.Artist, []models.Locations, []models.Dates, []models.Relation, error)`**
    *   **Description**: Fetches all primary data (artists, locations, dates, and relations) from the API concurrently to optimize loading times.

*   **`func GetMockArtists() []models.Artist`**
    *   **Description**: Provides a predefined slice of sample `Artist` data, typically used for testing or as a fallback when the external API is unavailable.

### Core Data Models (`models/models.go`)

Key data structures used for representing information within the application:

*   **`type Artist struct`**
    *   **Description**: Represents the basic information of a band or artist, including ID, image URL, name, members, creation date, and first album.

*   **`type Locations struct`**
    *   **Description**: Stores a list of all concert locations associated with a specific artist.

*   **`type Dates struct`**
    *   **Description**: Stores a list of all concert dates associated with a specific artist.

*   **`type Relation struct`**
    *   **Description**: Connects an artist's concert dates with their respective locations, mapping dates to events.

*   **`type ArtistDetails struct`**
    *   **Description**: Combines all relevant information for a single artist, including their basic details and a comprehensive list of all their concert dates and locations.

### Server Configuration (`server/server.go`, `server/types.go`)

*   **`type ServerConfig struct`**
    *   **Description**: Defines the configuration parameters for the HTTP server, such as address, multiplexer, and timeouts.

*   **`func NewServerConfig(Addr string, Mux *http.ServeMux, ReadTimeOut time.Duration, WriteTimeOut time.Duration, IdleTimeOut time.Duration, MaxHeaderBytes int) *ServerConfig`**
    *   **Description**: Initializes a new `ServerConfig` instance with specified network address, request multiplexer, and various timeout settings.

*   **`func (s *ServerConfig) Run()`**
    *   **Description**: Starts the HTTP server, making the application accessible via the configured network address.

## Contributing
Contributions are highly encouraged! If you have suggestions for improvements, new features, or bug fixes, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the [MIT License](LICENSE).
```