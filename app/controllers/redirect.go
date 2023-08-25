package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-chi/chi/v5"
	"goly.com/app/models"
)

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("./app/html/redirect.html")
	html, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	searchString := "{redirect}"
	goly := chi.URLParam(r, "goly")

	golyDb, err := models.GetByGolyUrl(goly)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	// Compile the regular expression
	re := regexp.MustCompile(searchString)

	// Replace matches using the regular expression
	redirect := re.ReplaceAllString(string(html), golyDb.Redirect)
	fmt.Println("Modified String:", html)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(redirect))
}

func SetupRedirect(app *chi.Mux) {
	r := chi.NewRouter()

	r.Get("/{goly}", handleRedirect)

	app.Mount("/", r)
}
