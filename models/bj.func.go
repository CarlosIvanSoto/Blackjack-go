package models

import (
	"fmt"
	"strings"
)

var (
	playerList = []Player{
		{100, "Carlos", 1000},
		{101, "Diana", 500},
		{102, "Roxer", 100000},
		{103, "Rosy", 500},
	}
)

func NewBlackjack() *Blackjack {
	deck := NewDeck()
	deck.Shuffle()

	game := Blackjack{
		Deck:       deck,
		IsGameOver: false,
	}

	return &game
}

// Methodo para iniciar el juego
func (g *Blackjack) Run() {
	fmt.Println("¡Juguemos al Blackjack!")
	// Repartir cartas a los jugadores
	g.IndexPlayerTurn = 0
	hands := make([]Hand, len(g.Players))
	g.Hands = hands
	for i := 0; i < len(g.Players); i++ {
		g.Deck.Deal(&g.Hands[i])
		g.Deck.Deal(&g.Hands[i])
	}
	// Repartir al dealer
	g.Deck.Deal(&g.DealerHand)
	g.Deck.Deal(&g.DealerHand)

	// Mostrar el juego de los jugadores
	for i := 0; i < len(g.Hands); i++ {
		fmt.Printf("Jugador %s tiene en su mano: %s\n", g.Players[i].Name, g.Hands[i].handString())
	}
	// Mostrar el juego del dealer
	fmt.Printf("Dealer tiene en su mano: %s\n", g.DealerHand.handString())

	more := g.checkPlayerStatus()
	if more {
		g.ask()
	}
}

// Pregunta al jugador si desea pedir o quedarse
func (g *Blackjack) ask() {
	i := g.IndexPlayerTurn
	fmt.Printf("Jugador %s: ¿Quieres pedir o quedarte? (p/q): ", g.Players[i].Name)
	var answer string
	fmt.Scan(&answer)
	if strings.ToLower(answer) == "p" {
		g.Hit()
	} else if strings.ToLower(answer) == "q" {
		g.Stay()
	} else {
		fmt.Println("Entrada inválida, inténtalo de nuevo")
		g.ask()
	}
}

// Agrega un nuevo jugador a la game con un id, para buscarlo
func (bj *Blackjack) AddPlayer(id int) {
	for i := range playerList {
		if playerList[i].Id == id {
			// Found!
			bj.Players = append(bj.Players, playerList[i])
			// bj.Scores = append(bj.Scores, playerList[i].Coins)
		}
	}
}

// Juega la partida del Dealer, hasta llegar a 17.
func (g *Blackjack) playDealer() {
	fmt.Printf("El dealer tiene: %s\n", g.DealerHand.handString())
	for g.DealerHand.handValue() < 17 {
		if g.DealerHand.handValue() == 21 {
			// Cobrar a todos, o dar empate
			fmt.Println("¡El dealer hizo Blackjack!")
			break
		}
		g.Deck.Deal(&g.DealerHand)
		fmt.Printf("El dealer tiene: %s\n", g.DealerHand.handString())
		if g.DealerHand.handValue() > 21 {
			fmt.Printf("¡El dealer se pasó con %d!\n", g.DealerHand.handValue())
			// O si perdio les paga a los que no se pasaron de 21
			break
		}
	}
	g.finishGame()
}
func (g *Blackjack) finishGame() {
	// Pagar a los que le ganaron al dealer
	g.IsGameOver = true
	// for i range g.Hands{

	// }

}

func (g *Blackjack) checkPlayerStatus() bool {
	i := g.IndexPlayerTurn
	if g.Hands[i].handValue() == 21 {
		fmt.Printf("¡Jugador %s hizo Blackjack!\n", g.Players[i].Name)
		// Dar la recompensa al jugador
		g.nextPlayer()
		return false
	} else if g.Hands[i].handValue() > 21 {
		fmt.Printf("¡El jugador %s se pasó con %d!\n", g.Players[i].Name, g.Hands[i].handValue())
		// Quitar creditos del jugador
		g.nextPlayer()
		return false
	}
	return true
}

// Pasa el turno al sig jugador
func (g *Blackjack) nextPlayer() {
	g.IndexPlayerTurn++
	// Verificar si sigue la jugada del dealer
	if len(g.Players) == g.IndexPlayerTurn {
		// turno del dealer
		g.playDealer()
	} else {
		g.ask()
	}
}

func (g *Blackjack) Hit() {
	i := g.IndexPlayerTurn
	g.Deck.Deal(&g.Hands[i])
	fmt.Printf("Jugador %s tiene en su mano: %s\n", g.Players[i].Name, g.Hands[i].handString())
	more := g.checkPlayerStatus()
	if more {
		g.ask()
	}
}

func (g *Blackjack) Stay() {
	i := g.IndexPlayerTurn
	fmt.Printf("Jugador %s se queda con %s\n", g.Players[i].Name, g.Hands[i].handString())
	g.nextPlayer()
	// g.checkPlayerStatus()
}
