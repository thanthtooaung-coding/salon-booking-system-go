package main

import (
	"log"
	"salon-booking-system-go/internal/db"
	"salon-booking-system-go/pkg/bookings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	bookings.RegisterRoutes(r, dbConn)

	r.Run(":8080")
}
