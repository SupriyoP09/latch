package booking

type MemoryStore struct {
	// seats -> bookings
	bookings map[string]Booking
}

// its a constructor for MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

// its a method to book a seat in the memory store
func (s *MemoryStore) Book(b Booking) error {
	// check if the seat is already booked
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil
}

// its a method to list all bookings for a specific movie in the memory store
func (s *MemoryStore) ListBooking(movieID string) []Booking {
	// creating a slice to hold the bookings for the specified movie
	var result []	Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
