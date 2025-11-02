package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParams(r)
	if err != nil {
		http.NotFound(w, r)
		// app.logger.Error(err.Error())
		return
	}

	fmt.Fprintf(w, "Show the detail of movie id: %d\n", id)
}
