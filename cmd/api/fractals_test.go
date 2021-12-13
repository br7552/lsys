package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/br7552/lsys/internal/data"
)

type input struct {
	Axiom      string            `json:"axiom"`
	Rules      map[string]string `json:"rules"`
	Depth      int               `json:"depth"`
	Angle      float64           `json:"angle"`
	StartAngle float64           `json:"start_angle"`
	Step       int               `json:"step"`
	Width      int               `json:"width"`
	Height     int               `json:"height"`
}

func TestGenerateFractalHandler(t *testing.T) {
	var app application

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	in := input{
		Axiom: "F++F++F",
		Rules: map[string]string{
			"F": "F-F++F-F",
		},
		Depth:      3,
		Angle:      60.0,
		StartAngle: 0.0,
		Step:       2,
		Width:      20,
		Height:     20,
	}

	js, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Post(ts.URL+"/v1/fractals", "", bytes.NewReader(js))
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, rs.StatusCode)
	}

	var env struct {
		Fractal data.Fractal `json="fractal"`
	}

	defer rs.Body.Close()
	dec := json.NewDecoder(rs.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&env)
	if err != nil {
		t.Fatal(err)
	}

	got := env.Fractal
	want := data.Fractal{
		Axiom:      in.Axiom,
		Rules:      in.Rules,
		Depth:      in.Depth,
		Angle:      in.Angle,
		StartAngle: in.StartAngle,
		Step:       in.Step,
		Width:      in.Width,
		Height:     in.Height,
	}

	data.Generate(&want)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
