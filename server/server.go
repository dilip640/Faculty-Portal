package server

import (
	"net/http"
	"os"

	"github.com/dilip640/Faculty-Portal/endpoints"
	"github.com/dilip640/Faculty-Portal/storage"
	log "github.com/sirupsen/logrus"
)

// Instance the struct for server
type Instance struct {
	httpServer *http.Server
}

// NewInstance just to get new instance of srver
func NewInstance() *Instance {
	s := &Instance{
		// just in case you need some setup here
	}

	return s
}

// Start starts the server
func (s *Instance) Start() { // Startup all dependencies
	storage.Initialize()
	endpoints.SetupRoutes()

	s.httpServer = &http.Server{Addr: os.Getenv("HOST_ADDRESS"), Handler: endpoints.Router}
	err := s.httpServer.ListenAndServe() // Blocks!
	if err != nil {
		log.WithError(err).Error("Http Server stopped unexpected")
	}
}
