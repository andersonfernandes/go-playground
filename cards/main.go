package main

func main() {
  // card := "Ace of spades"
  // var card string = "Ace of spades"

  // card := newCard()

  cards := deck{"Ace of Diamonds", newCard()}
  cards = append(cards, "Six of Spades")

  cards.print()
}

func newCard() string {
  return "Five of Diamonds"
}
