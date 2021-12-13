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

func TestgenerateFractalHandler(t *testing.T) {
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
		Step:       4,
		Width:      80,
		Height:     80,
	}

	js, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Post(ts.URL+"/fractals", "", bytes.NewReader(js))
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, rs.StatusCode)
	}

	var got data.Fractal

	defer rs.Body.Close()
	dec := json.NewDecoder(rs.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	expected := data.Fractal{
		Axiom:      in.Axiom,
		Rules:      in.Rules,
		Depth:      in.Depth,
		Angle:      in.Angle,
		StartAngle: in.StartAngle,
		Step:       in.Step,
		Width:      in.Width,
		Height:     in.Height,
	}

	data.Generate(&expected)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("reponse does not match expected")
	}
}
