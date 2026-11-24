package booking

// Booking represents a booking made by a user for a specific movie and seat.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

type BookingStore interface {
	Book(b Booking) (Booking, error)
	ListBooking(movieID string) []Booking
}
