package main

import (
	"net/http"

	"github.com/br7552/lsys/internal/data"
)

func (app *application) generateFractalHandler(w http.ResponseWriter,
	r *http.Request) {
	var input struct {
		Axiom      string            `json:"axiom"`
		Rules      map[string]string `json:"rules"`
		Depth      int               `json:"depth"`
		Angle      float64           `json:"angle"`
		StartAngle float64           `json:"start_angle"`
		Step       int               `json:"step"`
		Width      int               `json:"width"`
		Height     int               `json:"height"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// TODO validate input

	fractal := data.Fractal{
		Axiom:      input.Axiom,
		Rules:      input.Rules,
		Depth:      input.Depth,
		StartAngle: input.StartAngle,
		Step:       input.Step,
		Width:      input.Width,
		Height:     input.Height,
	}

	data.Generate(&fractal)

	err = app.writeJSON(w, http.StatusOK, envelope{"fractal": fractal}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
