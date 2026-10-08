package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

func main() {

	app := fiber.New()

	dbConc()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello Fiber")
	})

	log.Fatal(app.Listen(":8000"))
}

func dbConc() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat .env:", err)
	}

	var err error
	pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")

}
