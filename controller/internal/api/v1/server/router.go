package server

import (
	res "go-http/internal/api/response"

	"github.com/go-chi/chi/v5"
	"net/http"

	"go-http/internal/api/v1/modules/cron_module"
	"go-http/internal/api/v1/modules/exec_module"
)

func (s *server) router(app *chi.Mux) {

	// home_view
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res.Response(w, 200, nil, "Hello There!")
	})

	api := chi.NewRouter()
	apiPath := s.app.Conf().Backend.APIPath
	app.Mount(apiPath, api)

	cron_module.Router__(api)
	exec_module.Router__(api)
}
