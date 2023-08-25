package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"goly.com/app/models"
	"goly.com/app/utils"
)

type PostBody struct {
	Redirect string `json:"redirect"`
}

func getAllGolies(w http.ResponseWriter, r *http.Request) {
	golies, err := models.GetAll()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, err := json.Marshal(golies)
	w.Write(json)
}

func CreateGoly(w http.ResponseWriter, r *http.Request) {
	var body PostBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
		return
	}

	var goly models.Goly
	goly.Clicks = 0
	goly.Goly = utils.RandomURL(10)
	goly.Redirect = body.Redirect
	err = models.CreateGoly(goly)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func SetupGolies(api *chi.Mux) {
	r := chi.NewRouter()

	r.Get("/", getAllGolies)
	r.Post("/", CreateGoly)

	api.Mount("/goly", r)
}
