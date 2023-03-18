package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosIvanSoto/blackjack-go/models"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	// var game models.Blackjack
	game := models.NewBlackjack()
	// Crear un servidor Socket.IO
	io := socketio.NewServer(nil)

	// Manejar eventos de conexión de clientes
	io.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		game.AddPlayer(101)
		return nil

	})
	io.OnEvent("/", "start", func(s socketio.Conn) {
		fmt.Println("start for:", s.ID())
		game.Start(s)
	})

	io.OnEvent("/", "hit", func(s socketio.Conn) {
	})
	io.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Println("message:", msg)
		s.Emit("reply", "have "+msg)
	})
	io.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	io.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go io.Serve()
	defer io.Close()

	// Manejar rutas de HTTP
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// http.Handle("/socket.io/", handlers.CORS(originsOk, headersOk, methodsOk)(io))
	// http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(http.FileServer(http.Dir("./asset"))))
	//
	// http.Handle("/socket.io/", io)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     // Habilitar CORS
	//     w.Header().Set("Access-Control-Allow-Origin", "*")
	//     w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//     w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// 	http.FileServer(http.Dir("./asset"))
	//     // Resto del código aquí
	// })

	// Manejar rutas de HTTP
	http.Handle("/socket.io/", io)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	// Iniciar el servidor HTTP
	log.Println("Servidor escuchando en el puerto 8000...")
	log.Fatal(http.ListenAndServe(":7777", nil))

	// game := models.NewBlackjack()
	// game.AddPlayer(100)
	// game.AddPlayer(102)
	// game.Run()

}
