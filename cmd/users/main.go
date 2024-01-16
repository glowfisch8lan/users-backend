package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"idapp/users/cmd/users/common"
	"idapp/users/cmd/users/repository"
	"idapp/users/cmd/users/utils/env"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	log.Print("Init main app")

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Get the GITHUB_API_KEY environment variable
	dbPassword := env.Get("DB_PASSWORD")
	dbUser := env.Get("DB_USER")
	dbHost := env.Get("DB_HOST")
	dbPort := env.Get("DB_PORT")
	dbName := env.Get("DB_NAME")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dbUrl)

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("Connection db success...	")
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/users/create", func(w http.ResponseWriter, r *http.Request) {
		common.SendResponse(repository.CreateUser(), w, r)
	})

	http.ListenAndServe(":3000", r)
}
