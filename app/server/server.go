package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"goly.com/app/controllers"
)

func SetupAndListen() {
	app := chi.NewRouter()

	enroute(app)
	http.ListenAndServe(":3000", app)
}

func enroute(app *chi.Mux) {
	controllers.SetupGolies(app)
	controllers.SetupRedirect(app)
}
