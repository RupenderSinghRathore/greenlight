package main

import (
	"fmt"
	"net/http"

	"github.com/RupenderSinghRathore/greenlight/internal/data"
)

func (app application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "create a new movie")
	var input struct {
		Title   string   `json:"title"`
		Runtime int32    `json:"runtime"`
		Year    int32    `json:"year"`
		Genres  []string `json:"genres"`
	}
	err := readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
func (app application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	movie := data.Movie{
		ID:      id,
		Runtime: 120,
	}
	err = writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
