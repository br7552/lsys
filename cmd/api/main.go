package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const version = "1.0.0"

type config struct {
	port    int
	env     string
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	cfg config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development",
		"Environment (development|staging|production)")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2,
		"Rate limiter max requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4,
		"Rate limiter max burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true,
		"Enable rate limiter")

	flag.Parse()

	app := &application{
		cfg: cfg,
	}

	log.Printf("Starting server on port %d", app.cfg.port)

	err := app.serve()
	log.Fatal(err)
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv.ListenAndServe()
}
