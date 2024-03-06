

package main

import (
	"time"

	"github.com/google/uuid"
    "github.com/GnarlyLasagna/go-book-nook-rest-api/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Image     string    `json:"image"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Image:     dbUser.Image,
		APIKey:    dbUser.ApiKey,
	}
}

type Book struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"api_key"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseBookToBook(dbBook database.Book) Book {
	return Book{
		ID:        dbBook.ID,
		CreatedAt: dbBook.CreatedAt,
		UpdatedAt: dbBook.UpdatedAt,
		Name:      dbBook.Name,
		Url:       dbBook.Url,
		UserID:    dbBook.UserID,
	}
}

func databaseBooksToBooks(dbBooks []database.Book) []Book {
	books := []Book{}
	for _, dbBook := range dbBooks {
		books = append(books, databaseBookToBook(dbBook))
	}
	return books
}

type UserFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	BookID    uuid.UUID `json:"book_id"`
}

func databaseUserFollowToUserFollow(dbUserFollow database.UserFollow) UserFollow {
	return UserFollow{
		ID:        dbUserFollow.ID,
		CreatedAt: dbUserFollow.CreatedAt,
		UpdatedAt: dbUserFollow.UpdatedAt,
		UserID:    dbUserFollow.UserID,
		BookID:    dbUserFollow.BookID,
	}
}

func databaseUserFollowsToUserFollows(dbUserFollows []database.UserFollow) []UserFollow {
	userFollows := []UserFollow{}
	for _, dbUserFollow := range dbUserFollows {
		userFollows = append(userFollows, databaseUserFollowToUserFollow(dbUserFollow))
	}
	return userFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	BookID      uuid.UUID `json:"book_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		BookID:      dbPost.BookID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}
