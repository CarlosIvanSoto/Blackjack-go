package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
		suitStr = "♥️"
	case 2:
		suitStr = "♦️"
	case 3:
		suitStr = "♣️"
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

	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
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
	Hands        []Hand
	Scores       []int
	IsGameOver   bool
	IsPlayerTurn bool
}

func NewBlackjackGame(numPlayers int) *BlackjackGame {
	deck := NewDeck()
	deck.Shuffle()
	hands := make([]Hand, numPlayers)
	scores := make([]int, numPlayers)
	game := BlackjackGame{
		Deck:         deck,
		Hands:        hands,
		Scores:       scores,
		IsGameOver:   false,
		IsPlayerTurn: true,
	}
	for i := 0; i < numPlayers; i++ {
		game.Deck.Deal(&game.Hands[i])
		game.Deck.Deal(&game.Hands[i])
	}
	return &game
}

func (g *BlackjackGame) playDealer() {
	for g.SpotsLeft() > 0 {
		if g.Hands[g.SpotsLeft()-1].handValue() > 21 {
			continue
		}
		for g.Hands[g.SpotsLeft()-1].handValue() < 17 {
			g.Deck.Deal(&g.Hands[g.SpotsLeft()-1])
			if g.Hands[g.SpotsLeft()-1].handValue() > 21 {
				break
			}
		}
	}
	fmt.Println("El dealer tiene:", g.Hands[0].handString())
	for i := 1; i < len(g.Hands); i++ {
		if g.Hands[i].handValue() > 21 {
			fmt.Printf("El jugador %d se pasó con %d\n", i+1, g.Hands[i].handValue())
		} else if g.Hands[0].handValue() > 21 {
			fmt.Printf("El jugador %d ganó con %d\n", i+1, g.Hands[i].handValue())
			g.Scores[i]++
		} else if g.Hands[i].handValue() > g.Hands[0].handValue() {
			fmt.Printf("El jugador %d ganó con %d\n", i+1, g.Hands[i].handValue())
			g.Scores[i]++
		} else if g.Hands[i].handValue() < g.Hands[0].handValue() {
			fmt.Printf("El jugador %d perdió con %d\n", i+1, g.Hands[i].handValue())
			g.Scores[0]++
		} else {
			fmt.Printf("El jugador %d empató con el dealer con %d\n", i+1, g.Hands[i].handValue())
		}
	}
	g.IsGameOver = true
}

func (g *BlackjackGame) SpotsLeft() int {
	count := 0
	for i := range g.Hands {
		if len(g.Hands[i].Cards) == 0 {
			count++
		}
	}
	return len(g.Hands) - count
}

func (g *BlackjackGame) checkGameStatus() {
	spotsLeft := g.SpotsLeft()
	if spotsLeft == 0 {
		g.playDealer()
		return
	}
	if !g.IsPlayerTurn {
		g.playDealer()
		return
	}
	for i := 0; i < len(g.Hands); i++ {
		if g.Hands[i].handValue() == 21 {
			fmt.Printf("¡Jugador %d hizo Blackjack!\n", i+1)
			g.Scores[i]++
			g.Hands[i] = Hand{}
		} else if g.Hands[i].handValue() > 21 {
			fmt.Printf("¡El jugador %d se pasó con %d!\n", i+1, g.Hands[i].handValue())
			g.Scores[0]++
			g.Hands[i] = Hand{}
		}
	}
}

func (g *BlackjackGame) Hit(playerNum int) {
	if g.IsGameOver {
		return
	}
	g.Deck.Deal(&g.Hands[playerNum])
	fmt.Printf("Jugador %d tiene en su mano: %s\n", playerNum+1, g.Hands[playerNum].handString())
	g.checkGameStatus()
}

func (g *BlackjackGame) Stay(playerNum int) {
	if g.IsGameOver {
		return
	}
	fmt.Printf("Jugador %d se queda con %s\n", playerNum+1, g.Hands[playerNum].handString())
	g.checkGameStatus()
}

func (g *BlackjackGame) Run() {
	fmt.Println("¡Juguemos al Blackjack!")
	for i := 0; i < len(g.Hands); i++ {
		fmt.Printf("Jugador %d tiene en su mano: %s\n", i+1, g.Hands[i].handString())
	}
	for g.IsPlayerTurn {
		g.checkGameStatus()
		for i, _ := range g.Hands {
			if len(g.Hands[i].Cards) == 0 {
				continue
			}
			fmt.Printf("Jugador %d: ¿Quieres pedir o quedarte? (p/q): ", i+1)
			var answer string
			fmt.Scan(&answer)
			if strings.ToLower(answer) == "p" {
				g.Hit(i)
			} else if strings.ToLower(answer) == "q" {
				g.Stay(i)
			} else {
				fmt.Println("Entrada inválida, inténtalo de nuevo")
				i--
			}
		}
		if g.SpotsLeft() == 0 {
			break
		}
	}
	fmt.Println("\nPuntuaciones finales:")
	for i := range g.Hands {
		fmt.Printf("Jugador %d: %d\n", i+1, g.Scores[i])
	}
}

func main() {
	game := NewBlackjackGame(4)
	game.Run()
}
