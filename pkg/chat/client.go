package chat

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

// Client represents a WebSocket client connected to chat server
type Client struct {
	conn     *websocket.Conn
	room     *Room
	send     chan SocketMessage
	isClosed bool
}

func NewClient(conn *websocket.Conn, room *Room) *Client {
	return &Client{
		conn:     conn,
		room:     room,
		send:     make(chan SocketMessage),
		isClosed: false,
	}
}

func (c *Client) Read() {
	log.Println("Starting to read messages from client")
	defer func() {
		c.room.RemoveMember(c)
		c.conn.Close()
	}()

	for {
		var msg SocketMessage

		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected error: %v", err)
			} else {
				log.Printf("error reading message: %v", err)
			}

			break
		}
		// process the message
		log.Printf("Received message from client: %s\n", message)

		if err := json.Unmarshal(message, &msg); err == nil {
			c.room.Broadcast(c, msg)
		}

	}
	log.Println("Stopping reading messages from client")
}

func (c *Client) Write() {
	log.Println("Starting to write messages to client")
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The room closed the channel
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Printf("error writing close message: %v", err)
				}
				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				log.Printf("error writing JSON message: %v", err)
				return
			}
		}
	}
	log.Println("Stopping writing messages to client")
}
