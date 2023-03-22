package cards

type Card struct {
	Name              string   `json:"name"`
	Number            string   `json:"number"`
	Arcana            string   `json:"arcana"`
	Suit              string   `json:"suit"`
	Img               string   `json:"img"`
	Fortune           []string `json:"fortune"`
	Keywords          []string `json:"keywords"`
	Meanings          Position `json:"meanings"`
	Archetype         string   `json:"archetype"`
	HebrewAlphabet    string   `json:"hebrew alphabet"`
	Numerology        string   `json:"numerology"`
	Element           string   `json:"element"`
	MythicalSpiritual string   `json:"mythical/spiritual"`
	QuestionsToAsk    []string `json:"questions to ask"`
}

type Position struct {
	Upright  []string `json:"upright"`
	Reversed []string `json:"reversed"`
}
