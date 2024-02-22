package domain

import "github.com/gorilla/websocket"

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	chatRoom *ChatRoom
}

type ChatRoom struct {
	ID        string
	Members   map[*Client]bool
	Broadcast chan []byte
}

func NewChatRoom(id string) ChatRoom {
	return ChatRoom{
		ID:        id,
		Members:   make(map[*Client]bool),
		Broadcast: make(chan []byte),
	}
}
