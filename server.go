package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Map to store users and their connections
var clients = make(map[string]*websocket.Conn)
var mutex = sync.Mutex{}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	// Get username from query params
	username := r.URL.Query().Get("username")
	if username == "" {
		fmt.Println("Username not provided")
		return
	}

	// Store user connection
	mutex.Lock()
	clients[username] = conn
	mutex.Unlock()

	fmt.Printf("%s connected\n", username)

	// Listen for messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("%s disconnected\n", username)
			mutex.Lock()
			delete(clients, username)
			mutex.Unlock()
			break
		}

		fmt.Printf("Message from %s: %s\n", username, msg)
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	sender := r.URL.Query().Get("sender")
	receiver := r.URL.Query().Get("receiver")
	message := r.URL.Query().Get("message")

	mutex.Lock()
	conn, exists := clients[receiver]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Receiver not connected", http.StatusNotFound)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(sender+": "+message))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}

	w.Write([]byte("Message sent"))
}

func main() {
	http.HandleFunc("/ws", handleConnections) // WebSocket connection
	http.HandleFunc("/send", sendMessage)     // Message sending API

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
