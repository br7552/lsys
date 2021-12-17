package main

import (
	"net/http"

	"github.com/br7552/lsys/internal/router"
)

func (app *application) routes() http.Handler {
	mux := router.New()
	mux.NotFound = app.notFoundResponse
	mux.MethodNotAllowed = app.methodNotAllowedResponse

	mux.HandleFunc(http.MethodGet, "/v1/status", app.showStatusHandler)
	mux.HandleFunc(http.MethodPost, "/v1/fractals", app.generateFractalHandler)

	return app.recoverPanic(app.rateLimit(mux))
}
