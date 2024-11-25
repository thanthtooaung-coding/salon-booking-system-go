package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("SALON_BOOKING_SYSTEM_GO_DB_USER"),
		os.Getenv("SALON_BOOKING_SYSTEM_GO_DB_PASSWORD"),
		os.Getenv("SALON_BOOKING_SYSTEM_GO_DB_HOST"),
		os.Getenv("SALON_BOOKING_SYSTEM_GO_DB_PORT"),
		os.Getenv("SALON_BOOKING_SYSTEM_GO_DB_NAME"),
	)
	return sql.Open("mysql", dsn)
}
