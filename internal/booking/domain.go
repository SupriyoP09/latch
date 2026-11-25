package booking

import "errors"

var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
)

// Booking represents a booking made by a user for a specific movie and seat.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

type BookingStore interface {
	Book(b Booking) error
	ListBooking(movieID string) []Booking
}
