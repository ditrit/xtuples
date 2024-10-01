package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"

	"fmt"
	"go-http/pkg/app"
)

type server struct {
	app *app.App
}

func NewServer(app *app.App) *server {
	return &server{app}
}

// Start Starts the server on the specified host and port + with the defined routes
func (s *server) Start() {
	backend := chi.NewRouter()

	backend.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		middleware.ContentCharset("UTF-8", "Latin-1", ""),
		cors.Handler(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			MaxAge:         300,
		}),
	)

	s.router(backend) // use endpoints

	conf := s.app.Conf()
	address := fmt.Sprintf("%v:%v", conf.Backend.Host, conf.Backend.Port)
	http.ListenAndServe(address, backend)

}
