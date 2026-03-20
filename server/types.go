package server

import (
	"net/http"
	"time"
)

type ServerConfig struct {
	Addr           string
	Mux            *http.ServeMux
	ReadTimeOut    time.Duration
	WriteTimeOut   time.Duration
	IdleTimeOut    time.Duration
	MaxHeaderBytes int
}
