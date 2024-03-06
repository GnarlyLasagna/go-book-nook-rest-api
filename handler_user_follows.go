
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

    "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
    "github.com/GnarlyLasagna/go-book-nook-rest-api/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUserFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		BookID uuid.UUID `json:"book_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userFollow, err := apiCfg.DB.CreateUserFollow(r.Context(), database.CreateUserFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		BookID:    params.BookID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create book follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseUserFollowToUserFollow(userFollow))
}

func (apiCfg *apiConfig) handlerGetUserFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	userFollows, err := apiCfg.DB.GetUserFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get book follows: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseUserFollowsToUserFollows(userFollows))
}

func (apiCfg *apiConfig) handlerDeleteUserFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	userFollowIDStr := chi.URLParam(r, "userFollowID")
	userFollowID, err := uuid.Parse(userFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse book follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteUserFollow(r.Context(), database.DeleteUserFollowParams{
		ID:     userFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete book follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
