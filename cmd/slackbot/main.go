package main

import (
	"fmt"
	"log"
	"net/http"

	"slackbot/api"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

var (
	configPath string
)

const (
	defaultConfigPath = "./config.yaml"
)

type Config struct {
	Port     int    `env:"PORT" envDefault:"8080"`
	ClientID string `env:"CLIENT_ID"`
	Scopes   string `env:"SCOPES"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	mux := api.SetupRoutes()

	fmt.Printf("Starting web server at %d\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", cfg.Port), middleware.Recoverer(mux)); err != nil {
		log.Fatal(err)
	}
}
