package comms

import (
	"block-rogue/game/gamestate"
	"block-rogue/game/useractions"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if messageType == websocket.TextMessage {

			action := useractions.Action{}
			err = json.Unmarshal(msg, &action)
			if err != nil {
				fmt.Println("Error al deserializar mensaje:", err)
				continue
			}
			useractions.HandleUserAction(action)
		}
	}
}

func StartWebSocketServer(addr string, port int) {
	http.HandleFunc("/ws", wsHandler)

	fullAddr := fmt.Sprintf("%s:%d", addr, port)
	fmt.Printf("Servidor WebSocket en ws://%s/ws\n", fullAddr)

	err := http.ListenAndServe(fullAddr, nil)
	if err != nil {
		panic("Error al iniciar servidor: " + err.Error())
	}
}

func BroadcastState(state *gamestate.State) {
	payload, err := json.Marshal(state)
	if err != nil {
		fmt.Println("Error al serializar el estado:", err)
		return
	}

	mu.Lock()

	defer mu.Unlock()

	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
