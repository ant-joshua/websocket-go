package websocket

import (
	"fmt"
	"github.com/ant-joshua/dating-go/pkg/chat"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	chatService *chat.Service
	upgrader    websocket.Upgrader
}

func NewHandler(chatService *chat.Service) *Handler {
	return &Handler{
		chatService: chatService,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		log.Printf("Received: %s", message)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error writing message: %v", err)
			break
		}
	}
}

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		return
	}
	defer func() {
		log.Println("Closing WebSocket connection")
		conn.Close()
	}()

	log.Println("WebSocket connection established")

	// Assign the client to a chat based on query parameter
	chatID := r.URL.Query().Get("chat_id")
	fmt.Println("chatId", chatID)

	chatRoom := h.chatService.GetRoom(chatID)
	if chatRoom == nil {
		fmt.Println()
		chatRoom = h.chatService.CreateRoom(chatID)
		go chatRoom.Run() // Start the room's Run method
	}
	client := chat.NewClient(conn, chatRoom)
	chatRoom.AddMember(client)

	go client.Read()
	go client.Write()
}
