package models

import "fmt"

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
