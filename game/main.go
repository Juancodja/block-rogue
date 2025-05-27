package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al aceptar conexión:", err)
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("Nuevo cliente conectado")

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		fmt.Println("Cliente desconectado")
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

func broadcastLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		message := `{"x":50,"y":50,"color":"orange","width":40,"height":40}`

		mu.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				fmt.Println("Error al enviar mensaje:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
		mu.Unlock()
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	go broadcastLoop()

	fmt.Println("Servidor WebSocket en http://localhost:8888/ws")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic("Error al iniciar servidor: " + err.Error())
	}
}
