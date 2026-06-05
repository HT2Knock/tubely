package main

import (
	"encoding/json"
	"net/http"

	"github.com/T2Knock/tubely/internal/auth"
	"github.com/T2Knock/tubely/internal/database"
	"github.com/google/uuid"
)

func (s *Server) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Password == "" || params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required", nil)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	user, err := s.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Email:    params.Email,
		Password: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	user.Password = ""
	respondWithJSON(w, http.StatusCreated, user)
}
