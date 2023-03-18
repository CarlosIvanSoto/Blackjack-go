package models

import (
	"math/rand"
	"time"
)

// Crea un nuevo deck para su uso
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

// Mezcla el Deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Añade una carta del deck a la mano y la retira del deck
func (d *Deck) Deal(h *Hand) {
	if len(d.Cards) == 0 {
		return
	}
	card, cards := d.Cards[0], d.Cards[1:]
	h.Cards = append(h.Cards, card)
	d.Cards = cards
}
