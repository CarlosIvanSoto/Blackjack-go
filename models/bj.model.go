package models

type Player struct {
	Id    int
	Name  string
	Coins int
}
type Hand struct {
	Cards []Card
}
type Card struct {
	Rank int
	Suit int
}
type Deck struct {
	Cards []Card
}
type Blackjack struct {
	Message         string
	Deck            *Deck
	DealerHand      Hand
	IndexPlayerTurn int
	Players         []Player
	Hands           []Hand
	IsGameOver      bool
	Bet             []int
}
