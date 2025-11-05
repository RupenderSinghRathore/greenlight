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
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
