package cards

// Get returns one card, random or specific
func Get(name string, cards []*Card) ([]*Card, error) {
	var card []*Card

	for _, v := range cards {
		// if name is an empty string, return a random card
		// else
		if ToBase(v.Name) == name {
			card = append(card, v)
		}
	}

	return card, nil
}

// GetAll returns all cards
func GetAll() ([]*Card, error) {
	// create a slice of card images
	cards := []*Card{}

	// populate it with json
	// GET from database

	return cards, nil
}
