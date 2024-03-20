
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
    "github.com/GnarlyLasagna/go-book-nook-rest-api/internal/database"
)

func (apiCfg *apiConfig) handlerCreateBook(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Notes  string `json:"notes"`
		Image  string `json:"image"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	book, err := apiCfg.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create book: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseBookToBook(book))
}

func (apiCfg *apiConfig) handlerGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := apiCfg.DB.GetBooks(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get books: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseBooksToBooks(books))
}
