package main

import (
	"log"
	"net/http"
)

func (app *application) showStatusHandler(w http.ResponseWriter,
	r *http.Request) {
	status := map[string]string{
		"status":      "available",
		"environment": app.cfg.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"status": status}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}
