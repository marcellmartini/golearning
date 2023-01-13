package blackjack

// DeckOfCards is a map of all possible values of a deck of cards
var deckOfCards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// StrategyOptions map the possibles options of strategy
var strategyOptions = map[string]string{
	"Stand":            "S",
	"Hit":              "H",
	"Split":            "P",
	"AutomaticallyWin": "W",
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if value, ok := deckOfCards[card]; ok {
		return value
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myHand := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)

	switch {
	case myHand == 22:
		return strategyOptions["Split"]

	case myHand == 21:
		if dealerHand != 11 && dealerHand != 10 {
			return strategyOptions["AutomaticallyWin"]
		} else {
			return strategyOptions["Stand"]
		}

	case myHand >= 17 && myHand <= 20:
		return strategyOptions["Stand"]

	case myHand >= 12 && myHand <= 16:
		if dealerHand >= 7 {
			return strategyOptions["Hit"]
		} else {
			return strategyOptions["Stand"]
		}

	default:
		return strategyOptions["Hit"]
	}
}
