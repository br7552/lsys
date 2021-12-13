package main

import (
	"log"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request,
	status int, message interface{}) {
	err := app.writeJSON(w, status, envelope{"error": message}, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter,
	r *http.Request, err error) {
	log.Println(err)
	app.errorResponse(w, r, http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError))
}

func (app *application) notFoundResponse(w http.ResponseWriter,
	r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound,
		http.StatusText(http.StatusNotFound))
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter,
	r *http.Request) {
	app.errorResponse(w, r, http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed))
}

func (app *application) failedValidationResponse(w http.ResponseWriter,
	r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
