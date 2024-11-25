package bookings

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	group := router.Group("/api/v1/bookings")
	{
		group.GET("/:id", getBookingByID(db))
		group.GET("", listBookings(db))
		group.POST("", createBooking(db))
		group.PUT("/:id", updateBooking(db))
		group.DELETE("/:id", deleteBooking(db))
	}
}

func getBookingByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookingID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
			return
		}

		var booking Booking
		query := "SELECT id, user_id, salon_id, service_id, stylist_id FROM bookings WHERE id = ?"
		err = db.QueryRow(query, bookingID).Scan(
			&booking.ID, &booking.UserID, &booking.SalonID, &booking.ServiceID, &booking.StylistID,
		)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, booking)
	}
}

func listBookings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, user_id, salon_id, service_id, stylist_id FROM bookings")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var bookings []Booking
		for rows.Next() {
			var booking Booking
			err = rows.Scan(&booking.ID, &booking.UserID, &booking.SalonID, &booking.ServiceID, &booking.StylistID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			bookings = append(bookings, booking)
		}

		c.JSON(http.StatusOK, bookings)
	}
}

func createBooking(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBooking Booking
		if err := c.ShouldBindJSON(&newBooking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO bookings (user_id, salon_id, service_id, stylist_id) VALUES (?, ?, ?, ?)"
		result, err := db.Exec(query, newBooking.UserID, newBooking.SalonID, newBooking.ServiceID, newBooking.StylistID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		insertedID, _ := result.LastInsertId()
		newBooking.ID = int(insertedID)

		c.JSON(http.StatusCreated, newBooking)
	}
}

func updateBooking(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookingID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
			return
		}

		var updatedBooking Booking
		if err := c.ShouldBindJSON(&updatedBooking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "UPDATE bookings SET user_id = ?, salon_id = ?, service_id = ?, stylist_id = ? WHERE id = ?"
		_, err = db.Exec(query, updatedBooking.UserID, updatedBooking.SalonID, updatedBooking.ServiceID, updatedBooking.StylistID, bookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		updatedBooking.ID = bookingID
		c.JSON(http.StatusOK, updatedBooking)
	}
}

func deleteBooking(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookingID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
			return
		}

		query := "DELETE FROM bookings WHERE id = ?"
		_, err = db.Exec(query, bookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
	}
}
