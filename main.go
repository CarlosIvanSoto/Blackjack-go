package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	game := NewBlackjackGame()
	game.Run()
}

type Card struct {
	Rank int
	Suit int
}

func (c Card) String() string {
	var suitStr string
	switch c.Suit {
	case 0:
		suitStr = "♠"
	case 1:
		suitStr = "♥"
	case 2:
		suitStr = "♦"
	case 3:
		suitStr = "♣"
	}
	var rankStr string
	switch c.Rank {
	case 1:
		rankStr = "A"
	case 11:
		rankStr = "J"
	case 12:
		rankStr = "Q"
	case 13:
		rankStr = "K"
	default:
		rankStr = fmt.Sprintf("%d", c.Rank)
	}
	return fmt.Sprintf("%s%s", rankStr, suitStr)
}

type Hand struct {
	Cards []Card
}

func (h Hand) handValue() int {
	value := 0
	numAces := 0
	for _, c := range h.Cards {
		if c.Rank == 1 {
			numAces++
			value += 11
		} else if c.Rank >= 10 {
			value += 10
		} else {
			value += c.Rank
		}
	}
	for numAces > 0 && value > 21 {
		value -= 10
		numAces--
	}
	return value
}

func (h Hand) handString() string {
	cardStrings := []string{}
	for _, c := range h.Cards {
		cardStrings = append(cardStrings, c.String())
	}
	s := strings.Join(cardStrings, " | ")
	s += fmt.Sprintf(" - Total: %d", h.handValue())
	return s
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	for i := 1; i <= 13; i++ {
		for j := 0; j < 4; j++ {
			card := Card{Rank: i, Suit: j}
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	// Recorre la baraja en orden inverso y para cada carta, selecciona una carta aleatoria anterior y las intercambia.
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

	// for i := range d.Cards {
	// 	j := rand.Intn(i + 1)
	// 	d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	// }
}

func (d *Deck) Deal(h *Hand) {
	if len(d.Cards) == 0 {
		return
	}
	card, cards := d.Cards[0], d.Cards[1:]
	h.Cards = append(h.Cards, card)
	d.Cards = cards
}

type BlackjackGame struct {
	Deck         *Deck
	PlayerHand   Hand
	DealerHand   Hand
	PlayerScore  int
	DealerScore  int
	IsGameOver   bool
	IsPlayerTurn bool
}

func NewBlackjackGame() *BlackjackGame {
	deck := NewDeck()
	deck.Shuffle()
	playerHand := Hand{}
	dealerHand := Hand{}
	game := BlackjackGame{
		Deck:         deck,
		PlayerHand:   playerHand,
		DealerHand:   dealerHand,
		PlayerScore:  0,
		DealerScore:  0,
		IsGameOver:   false,
		IsPlayerTurn: true,
	}
	game.Deck.Deal(&game.PlayerHand)
	game.Deck.Deal(&game.DealerHand)
	game.Deck.Deal(&game.PlayerHand)
	game.Deck.Deal(&game.DealerHand)
	return &game
}

func (g *BlackjackGame) playDealer() {
	if g.PlayerHand.handValue() > 21 {
		fmt.Println("Dealer:", g.DealerHand.handString())
		g.IsGameOver = true
		return
	}
	for g.DealerHand.handValue() < 17 {
		g.Deck.Deal(&g.DealerHand)
		// fmt.Println("Dealer:", g.DealerHand.handString())
	}
	fmt.Println("Dealer:", g.DealerHand.handString())
	g.DealerScore = g.DealerHand.handValue()
}

func (g *BlackjackGame) checkGameStatus() {
	if g.PlayerHand.handValue() > 21 {
		fmt.Println("Te pasaste! El Dealer Gano!")
		g.IsGameOver = true
		return
	}
	if g.DealerHand.handValue() > 21 {
		fmt.Println("El Dealer se paso! Tu Ganaste!")
		g.IsGameOver = true
		return
	}
	if !g.IsPlayerTurn {
		if g.PlayerHand.handValue() > g.DealerHand.handValue() {
			fmt.Println("Tu Ganaste!")
			g.IsGameOver = true
			return
		} else if g.DealerHand.handValue() > g.PlayerHand.handValue() {
			fmt.Println("El Dealer Gano!")
			g.IsGameOver = true
			return
		} else {
			fmt.Println("Empate, Nadie Gano!")
			g.IsGameOver = true
			return
		}
	}
}

func (g *BlackjackGame) Hit() {
	if g.IsGameOver {
		return
	}
	g.Deck.Deal(&g.PlayerHand)
	if g.PlayerHand.handValue() >= 21 {
		//g.IsGameOver = true
		g.Stay()
	} else {
		fmt.Println("Tienes en tu mano:", g.PlayerHand.handString())
		fmt.Print("¿Quieres pedir de nuevo? (y/n): ")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) == "y" {
			g.Hit()
		} else {
			g.Stay()
		}
	}
}

func (g *BlackjackGame) Stay() {
	if g.IsGameOver {
		return
	}
	fmt.Println("Te quedaste con:", g.PlayerHand.handString())
	g.IsPlayerTurn = false
	g.playDealer()
	g.checkGameStatus()
}

func (g *BlackjackGame) Run() {
	fmt.Println("Juguemos al blackjack!")
	fmt.Println("Tienes en tu mano:", g.PlayerHand.handString())
	fmt.Println("Tarjeta boca arriba del Dealer:", g.DealerHand.Cards[0].String())
	for g.IsPlayerTurn {
		if g.PlayerHand.handValue() == 21 {
			fmt.Println("Hiciste ¡Blackjack! ¡Ganaste!")
			break
		}
		if g.DealerHand.handValue() == 21 {
			fmt.Println("Dealer:", g.DealerHand.handString())
			fmt.Println("El Dealer hizo ¡Blackjack! Perdiste!")
			break
		}
		fmt.Print("¿Quieres pedir o te quedas? pedir(hit)/quedarse(stay) (h/s): ")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) == "h" {
			g.Hit()
		} else if strings.ToLower(answer) == "s" {
			g.Stay()
		} else {
			fmt.Println("Entrada no válida, inténtalo de nuevo.")
		}
	}
}
