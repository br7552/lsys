package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/br7552/lsys/internal/data"
)

func (app *application) generateFractalHandler(w http.ResponseWriter,
	r *http.Request) {
	/*
		fractal := data.Fractal{
			Axiom:      "F++F++F",
			Rules:      []string{"F", "F-F++F-F"},
			Depth:      3,
			Angle:      60.0,
			StartAngle: 0.0,
			Step:       2,
			Width:      60,
			Height:     90,
		}
	*/

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

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&fractal)
	if err != nil {
		log.Println(err)
		return
	}

	data.Generate(&fractal)

	err = app.writeJSON(w, http.StatusOK, envelope{"fractal": fractal}, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}
