package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	Participants []primitive.ObjectID `bson:"participants"`
}

type ChatMessage struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	ChatID  primitive.ObjectID `bson:"chat_id"`
	UserID  primitive.ObjectID `bson:"user_id"`
	Content string             `bson:"content"`
	Type    string             `bson:"type"` // text, image, video, audio, file
	ReplyTo primitive.ObjectID `bson:"reply_to,omitempty"`
	SentAt  int64              `bson:"sent_at"`
}

type WebSocketMessage struct {
	Type string                 `json:"type"` // message, user_joined, user_left
	Data map[string]interface{} `json:"data"`
}
