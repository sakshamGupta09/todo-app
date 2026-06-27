package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"todo-app/internal/app"
	"todo-app/internal/database"
	"todo-app/internal/todos"

	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("Database connection string not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	db, err := database.NewPostgresPool(databaseUrl, ctx)
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}
	fmt.Println("Connected to DB ....")

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	application := &app.App{
		DB:      db,
		Decoder: decoder,
	}

	mux := http.NewServeMux()

	todos.Setup(mux, application)

	log.Fatal(http.ListenAndServe(":80", mux))
}
