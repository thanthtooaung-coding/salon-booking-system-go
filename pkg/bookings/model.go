package bookings

type Booking struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	SalonID   int `json:"salonId"`
	ServiceID int `json:"serviceId"`
	StylistID int `json:"stylistId"`
}
