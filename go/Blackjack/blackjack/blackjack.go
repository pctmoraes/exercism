package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValue int

	switch card {
	case "ace":
		cardValue = 11
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten":
		cardValue = 10
	case "jack":
		cardValue = 10
	case "queen":
		cardValue = 10
	case "king":
		cardValue = 10
	default:
		cardValue = 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var option string

	parsedCard1 := ParseCard(card1)
	parsedCard2 := ParseCard(card2)
	parsedDealerCard := ParseCard(dealerCard)

	sum := parsedCard1 + parsedCard2

	switch {
	case card1 == "ace" && card2 == "ace":
		option = "P"
	case sum <= 11:
		option = "H"
	case sum  >= 12 && sum <= 16 && parsedDealerCard >= 7:
		option = "H"
	case sum  >= 12 && sum <= 16 && parsedDealerCard < 7:
		option = "S"
	case sum >= 17 && sum <= 20:
		option = "S"
	case sum == 21 && parsedDealerCard < 10:
		option = "W"
	default:
		option = "S"
	}

	return option
}
