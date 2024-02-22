package main

import (
	"github.com/ant-joshua/dating-go/pkg/chat"
	"github.com/ant-joshua/dating-go/pkg/websocket"
	"log"
	"net/http"
)

func main() {

	chatService := chat.NewService()

	wsHandler := websocket.NewHandler(chatService)

	http.HandleFunc("/ws", wsHandler.HandleConnections)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
