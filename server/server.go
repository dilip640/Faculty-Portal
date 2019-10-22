package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dilip640/Faculty-Portal/endpoints"
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
	endpoints.SetupRoutes()
	s.httpServer = &http.Server{Addr: os.Getenv("HOST_ADDRESS"), Handler: endpoints.Router}
	err := s.httpServer.ListenAndServe() // Blocks!
	if err != http.ErrServerClosed {
		log.WithError(err).Error("Http Server stopped unexpected")
		s.Shutdown()
	} else {
		log.WithError(err).Info("Http Server stopped")
	}
}

// Shutdown shutdown the server
func (s *Instance) Shutdown() {
	if s.httpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			log.WithError(err).Error("Failed to shutdown http server gracefully")
		} else {
			s.httpServer = nil
		}
	}
}
