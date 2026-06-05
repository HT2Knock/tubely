package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	DBPath           string `env:"DB_PATH"`
	JWTSecret        string `env:"JWT_SECRET"`
	Platform         string `env:"PLATFORM"`
	FilePathRoot     string `env:"FILEPATH_ROOT"`
	AssetsRoot       string `env:"ASSETS_ROOT"`
	S3Bucket         string `env:"S3_BUCKET"`
	S3Region         string `env:"S3_REGION"`
	S3CFDistribution string `env:"S3_CF_DISTRIBUTION"`
	Port             string `env:"PORT"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Missing .env file")
	}

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Failed to parse env vars")
	}

	err := cfg.ensureAssetsDir()
	if err != nil {
		log.Fatalf("Couldn't create assets directory: %v", err)
	}

	server := NewServer(&cfg)

	log.Printf("Serving on: http://localhost:%s/app/\n", cfg.Port)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
