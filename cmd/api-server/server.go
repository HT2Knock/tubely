package main

import (
	"log"
	"net/http"
	"time"

	"github.com/T2Knock/tubely/internal/database"
	"github.com/google/uuid"
)

var videoThumbnails = map[uuid.UUID]thumbnail{}

type thumbnail struct {
	mediaType string
	data      []byte
}

type Server struct {
	cfg *config
	db  database.Client
}

func NewServer(cfg *config) *http.Server {
	db, err := database.NewClient(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		cfg: cfg,
		db:  db,
	}

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      s.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
