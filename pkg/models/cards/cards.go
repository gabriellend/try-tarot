package cards

type Card struct {
	Name             string              `json:"name"`
	Suit             string              `json:"suit"`
	Element          string              `json:"element"`
	IsMinor          bool                `json:"isMinor"` // false means it's Major
	IsCourt          bool                `json:"isCourt"`
	IsUpright        bool                `json:"isUpright"` // false mean it's reversed
	UprightMeanings  []string            `json:"uprightMeanings"`
	ReversedMeanings []string            `json:"reversedMeanings"`
	RelatedCards     []string            `json:"relatedCards"`
	Symbols          map[string][]string `json:"symbols"`
}
