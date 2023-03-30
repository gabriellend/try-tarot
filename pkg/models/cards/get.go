package cards

import (
	"math/rand"
	"strings"
)

// Get returns the named card
func Get(name string, cards []*Card) ([]*Card, error) {
	var card []*Card

	for _, v := range cards {
		if ToBase(v.Name) == name {
			card = append(card, v)
		}
	}

	return card, nil
}

// GetSet returns a particular set of the cards (i.e. "Major")
func GetSet(set string, cards []*Card) ([]*Card, error) {
	var flashCards []*Card

	for _, v := range cards {
		switch {
		case strings.ToLower(v.Arcana) == set:
			flashCards = append(flashCards, v)
		case strings.ToLower(v.Suit) == set:
			flashCards = append(flashCards, v)
		case set == "all":
			flashCards = append(flashCards, v)
		}
	}

	return flashCards, nil
}

// Random returns a random card from a specified set (i.e. "Major")
func Random(cards []*Card) Card {
	var card *Card
	major := len(cards) == 22
	minor := len(cards) == 56
	suit := len(cards) == 14

	switch {
	case major:
		card = cards[rand.Intn(22)]
	case minor:
		card = cards[rand.Intn(56)]
	case suit:
		card = cards[rand.Intn(14)]
	default:
		card = cards[rand.Intn(78)]
	}

	return *card
}
