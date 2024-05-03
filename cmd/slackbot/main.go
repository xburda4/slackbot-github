package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"slackbot/api"
	"slackbot/ent"
	"slackbot/service"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	database, err := ent.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer database.Close()
	// Run the auto migration tool.
	if err := database.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s, err := service.NewService(database)
	if err != nil {
		log.Fatal(err)
	}

	h := api.NewHandler(s)

	fmt.Printf("Starting web server at %d\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", cfg.Port), middleware.Recoverer(h.Mux)); err != nil {
		log.Fatal(err)
	}
}
