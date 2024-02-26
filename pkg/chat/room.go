package chat

import (
	"log"
	"sync"
)

type Room struct {
	ID       string
	Members  map[*Client]bool
	register chan *Client
	mutex    sync.RWMutex
}

func NewRoom(id string) *Room {
	return &Room{
		ID:       id,
		Members:  make(map[*Client]bool),
		register: make(chan *Client),
	}
}

func (r *Room) Run() {
	log.Printf("Room Run")
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				log.Println("masuk sini")
				return // Exit the loop if the channel is closed
			}
			r.mutex.Lock()
			r.Members[client] = true
			r.mutex.Unlock()
			// Additional select cases for handling messages, etc.
		}
	}
}

func (r *Room) AddMember(client *Client) {
	log.Println("Adding member to room:", r.ID)
	r.register <- client
}

func (r *Room) RemoveMember(client *Client) {
	log.Println("Removing member from room:", r.ID)
	r.mutex.Lock()
	delete(r.Members, client)
	r.mutex.Unlock()
}

// Broadcast sends a message to all members of the chat room
func (r *Room) Broadcast(sender *Client, message SocketMessage) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	//var toRemove []*Client
	for member := range r.Members {
		if member == sender {
			continue // Skip the sender
		}
		select {
		case member.send <- message:
		default:
			close(member.send)
			delete(r.Members, member)

		}
	}
}
