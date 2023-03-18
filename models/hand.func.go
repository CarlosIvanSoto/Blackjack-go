package models

import (
	"fmt"
	"strings"
)

func (h Hand) handString() string {
	cardStrings := []string{}
	for _, c := range h.Cards {
		cardStrings = append(cardStrings, c.String())
	}
	s := strings.Join(cardStrings, " | ")
	s += fmt.Sprintf(" - Total: %d", h.handValue())
	return s
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
