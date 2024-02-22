package chat

import (
	"log"
	"sync"
)

type Service struct {
	rooms map[string]*Room
	mutex sync.RWMutex
}

func NewService() *Service {
	return &Service{
		rooms: make(map[string]*Room),
	}
}

// GetRoom returns a chat room by its ID. If the room doesn't exist, it returns nil.
func (s *Service) GetRoom(roomID string) *Room {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.rooms[roomID]
}

// CreateRoom creates a new chat room with the given ID and returns it.
// If a room with the same ID already exists, it returns the existing room.
func (s *Service) CreateRoom(roomID string) *Room {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if room, exists := s.rooms[roomID]; exists {
		return room
	}
	room := NewRoom(roomID)
	s.rooms[roomID] = room
	return room
}

func (s *Service) DeleteRoom(roomID string) {
	log.Printf("deleting room %s", roomID)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.rooms, roomID)

}
