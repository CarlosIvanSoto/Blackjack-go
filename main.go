package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosIvanSoto/blackjack-go/models"
	socketio "github.com/googollee/go-socket.io"
)

const (
	port = "8000"
)

func CorsMiddleware(handler http.Handler, allowOrigin string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	// var game models.Blackjack
	game := models.NewBlackjack()
	// Crear un servidor Socket.IO
	server := socketio.NewServer(nil)

	// Manejar eventos de conexión de clientes
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		game.AddPlayer(101)
		return nil

	})
	server.OnEvent("/", "start", func(s socketio.Conn) {
		fmt.Println("start for:", s.ID())
		game.Start(s)
	})
	server.OnEvent("/", "hit", func(s socketio.Conn) {
	})
	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Println("message:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		game.RemovePlayer(101)
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", CorsMiddleware(server, "http://127.0.0.1:5500"))
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Servidor escuchando en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// game := models.NewBlackjack()
	// game.AddPlayer(100)
	// game.AddPlayer(102)
	// game.Run()

}
