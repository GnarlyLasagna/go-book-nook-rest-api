package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"fmt"
//    "time"

	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
    "github.com/go-chi/chi/v5"
    "github.com/GnarlyLasagna/go-book-nook-rest-api/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found")
	}

    dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database URL is not found in the environment")
	}

    conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

    v1Router.Post("/users", apiCfg.handlerCreateUser)
    v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

    v1Router.Post("/books", apiCfg.middlewareAuth(apiCfg.handlerCreateBook))
	v1Router.Get("/books", apiCfg.handlerGetBooks)

    v1Router.Post("/user_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateUserFollow))
	v1Router.Get("/user_follows", apiCfg.middlewareAuth(apiCfg.handlerGetUserFollows))
    v1Router.Delete("/user_follows/{userFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteUserFollow))

//    v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))


	router.Mount("/v1", v1Router)	

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("server starting on %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}
