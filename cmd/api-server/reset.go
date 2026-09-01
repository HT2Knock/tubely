package main

import (
	"log"
	"net/http"
)

func (s *Server) handlerReset(w http.ResponseWriter, r *http.Request) {
	if s.cfg.Platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		if _, err := w.Write([]byte("Reset is only allowed in dev environment.")); err != nil {
			log.Printf("failed to write response: %v", err)
		}
		return
	}

	err := s.db.Reset()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't reset database", err)
		return
	}
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("Database reset to initial state")); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
