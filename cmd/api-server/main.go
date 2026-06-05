package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := godotenv.Load(".env"); err != nil {
		log.Warn().Err(err).Msg("")
	}

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed to parse env vars")
	}

	err := cfg.ensureAssetsDir()
	if err != nil {
		log.Error().Err(err).Msg("Couldn't create assets directory")
	}

	server := NewServer(&cfg)

	log.Info().Str("port", cfg.Port).Msg("Serving on: http://localhost:%s/app/")
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Panic().Msg(fmt.Sprintf("http server error: %s", err))
	}
}
