package main

import (
	"fmt"
	"log"
	"net/http"

	"slackbot/slack"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"
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

	/*workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	http.Handle("/", filesDir)*/

	r := chi.NewRouter()

	r.Route("/slack", func(r chi.Router) {
		r.Post("/command", slack.HandleCommand)
	})

	fmt.Printf("Starting web server at %d\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), middleware.Recoverer(r)); err != nil {
		log.Fatal(err)
	}
}
