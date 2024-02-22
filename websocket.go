package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]map[string]interface{}) // connected clients
var broadcast = make(chan WebSocketMessage)                    // broadcast channel

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Assign the client to a chat based on query parameter
	chatID := r.URL.Query().Get("chat_id")
	userID := r.URL.Query().Get("user_id")
	clients[conn] = map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	for {
		var msg WebSocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}
		//msg.ChatID, err = primitive.ObjectIDFromHex(chatID)
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}
		broadcast <- msg // broadcast the message
	}
}

func saveMessage(message ChatMessage) error {
	message.SentAt = time.Now().Unix()
	_, err := messagesCollection.InsertOne(context.Background(), message)
	return err
}

func getMessages(chatID string) ([]ChatMessage, error) {
	var messages []ChatMessage
	filter := bson.M{"chat_id": chatID}
	cursor, err := messagesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message ChatMessage
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func handleMessages() {
	for {
		msg := <-broadcast // grab the next message from the broadcast channel

		// Save the message to the database

		messageData := msg.Data

		//err := saveMessage(messageData)
		//if err != nil {
		//	log.Println(err)
		//	continue
		//}

		// broadcast the message to all clients
		for client, socket := range clients {

			messageChatID := messageData["chat_id"].(string)

			if socket["chat_id"] == messageChatID && messageData["user_id"] != clients[client]["user_id"] {
				err := client.WriteJSON(msg)

				fmt.Println("Chat ID: ", socket["chat_id"])
				fmt.Println("User ID: ", messageData["user_id"])

				if err != nil {
					log.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
