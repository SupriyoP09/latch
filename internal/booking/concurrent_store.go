package booking

import "sync"

// ConcurentStore is a thread-safe implementation of BookingStore that allows concurrent access to the bookings map.
type ConcurentStore struct {
	bookings     map[string]Booking
	sync.RWMutex // RWMutex to allow concurrent reads and exclusive writes
}

func NewConcurentStore() *ConcurentStore {
	return &ConcurentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurentStore) Book(b Booking) error {
	// Lock and defer unlock to ensure thread safety when accessing the bookings map
	s.Lock()
	defer s.Unlock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil
}

func (s *ConcurentStore) ListBooking(movieID string) []Booking {
	// Lock for reading and defer unlock to ensure thread safety when accessing the bookings map
	s.RLock()
	defer s.RUnlock()
	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
