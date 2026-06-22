package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"todo-app/internal/database"

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
	fmt.Println("Connected to DB ....", db)

}
