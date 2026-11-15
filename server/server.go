package server

import (
	"log"
	"net/http"

	"restapi/handlers"
)

// Config holds the listen and file settings for the Server.
type Config struct {
	HTTPAddr string // e.g. ":5000"

	// HTTPSAddr, CertFile and KeyFile enable the HTTPS listener.
	// Leave HTTPSAddr empty to run HTTP only (e.g. on Render, which provides HTTPS itself).
	HTTPSAddr string
	CertFile  string
	KeyFile   string

	// StaticDir is the built frontend (frontend/dist). When set, the server
	// serves the frontend at / and the API also under /api. Leave empty to
	// serve only the API.
	StaticDir string
}

// Server holds the router and listen settings.
type Server struct {
	mux *http.ServeMux
	cfg Config
}

// New creates a Server and registers all routes.
func New(cfg Config) *Server {
	s := &Server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	api := http.NewServeMux()

	vehicles := handlers.NewVehiclesHandler()
	api.HandleFunc("GET /vehicles", vehicles.Get)
	api.HandleFunc("POST /vehicles", vehicles.Create)

	docs := handlers.NewDocsHandler()
	api.HandleFunc("GET /docs", docs.UI)
	api.HandleFunc("GET /docs/openapi.json", docs.Spec)

	// The API stays at the root paths (/vehicles, /docs) as before.
	s.mux.Handle("/vehicles", api)
	s.mux.Handle("/docs", api)
	s.mux.Handle("/docs/", api)

	if s.cfg.StaticDir != "" {
		// The frontend calls /api/vehicles, the same path the Vite dev server forwards.
		s.mux.Handle("/api/", http.StripPrefix("/api", api))
		s.mux.Handle("/", newSPAHandler(s.cfg.StaticDir))
		log.Printf("Serving frontend from %s", s.cfg.StaticDir)
	}
}

// Start runs the HTTP listener, plus the HTTPS listener when it is configured.
// It blocks until one of them fails.
func (s *Server) Start() error {
	errCh := make(chan error, 2)

	go func() {
		log.Printf("HTTP listening on %s", s.cfg.HTTPAddr)
		errCh <- http.ListenAndServe(s.cfg.HTTPAddr, s.mux)
	}()

	if s.cfg.HTTPSAddr != "" {
		go func() {
			log.Printf("HTTPS listening on %s", s.cfg.HTTPSAddr)
			errCh <- http.ListenAndServeTLS(s.cfg.HTTPSAddr, s.cfg.CertFile, s.cfg.KeyFile, s.mux)
		}()
	}

	return <-errCh
}
