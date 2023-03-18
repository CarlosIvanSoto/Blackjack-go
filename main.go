package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CarlosIvanSoto/blackjack-go/game"
	"github.com/CarlosIvanSoto/blackjack-go/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Action string `json:"action"`
	Player int    `json:"player"`
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	game := game.NewBlackjackGame(4)
	game.BroadcastGameState(conn)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer el mensaje:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Error al decodificar el mensaje:", err)
			break
		}

		switch msg.Action {
		case "hit":
			game.Hit(msg.Player)
		case "stay":
			game.Stay(msg.Player)
		}

		game.BroadcastGameState(conn)
	}
}

func main() {
	// http.HandleFunc("/", handleConnection)

	// log.Println("Escuchando en :8080...")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	game := models.NewBlackjack()
	game.AddPlayer(100)
	game.AddPlayer(102)
	game.Run()

}

// Agrega estas funciones a la estructura BlackjackGame
