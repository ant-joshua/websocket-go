package main

import (
	"log"
	"net/http"
)

func main() {

	InitMongo()
	// work with the connection

	http.HandleFunc("/ws", HandleWebsocket)

	go handleMessages()

	log.Println("Server started on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
