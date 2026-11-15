package server

import (
	"log"
	"net/http"

	"restapi/handlers"
)

// Server holds the router and listen settings.
type Server struct {
	mux       *http.ServeMux
	httpAddr  string
	httpsAddr string
	certFile  string
	keyFile   string
}

// New creates a Server and registers all routes.
func New(httpAddr, httpsAddr, certFile, keyFile string) *Server {
	s := &Server{
		mux:       http.NewServeMux(),
		httpAddr:  httpAddr,
		httpsAddr: httpsAddr,
		certFile:  certFile,
		keyFile:   keyFile,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	vehicles := handlers.NewVehiclesHandler()
	s.mux.HandleFunc("GET /vehicles", vehicles.Get)
	s.mux.HandleFunc("POST /vehicles", vehicles.Create)

	docs := handlers.NewDocsHandler()
	s.mux.HandleFunc("GET /docs", docs.UI)
	s.mux.HandleFunc("GET /docs/openapi.json", docs.Spec)
}

// Start runs the HTTP and HTTPS listeners. It blocks until one of them fails.
func (s *Server) Start() error {
	errCh := make(chan error, 2)

	go func() {
		log.Printf("HTTP listening on %s", s.httpAddr)
		errCh <- http.ListenAndServe(s.httpAddr, s.mux)
	}()

	go func() {
		log.Printf("HTTPS listening on %s", s.httpsAddr)
		errCh <- http.ListenAndServeTLS(s.httpsAddr, s.certFile, s.keyFile, s.mux)
	}()

	return <-errCh
}
