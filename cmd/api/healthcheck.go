package main

import (
	"net/http"
)

// type

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// js := fmt.Sprintf(
	// 	`{"status":"available", "environment":%q, "version":%q}`,
	// 	app.config.env,
	// 	version,
	// )
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			"server encountered a problem and could't process your response",
			http.StatusInternalServerError,
		)
	}
}
