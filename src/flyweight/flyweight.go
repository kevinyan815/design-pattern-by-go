package main

import "fmt"

var pokerCards = map[int]*Card{
	1: {
		Name:  "A",
		Color: "紅",
	},
	2: {
		Name:  "A",
		Color: "黑",
	},
	// 其他卡牌
}

type Card struct {
	Name  string
	Color string
}

type PokerGame struct {
	Cards map[int]*Card
}

func NewPokerGame() *PokerGame {
	board := &PokerGame{Cards: map[int]*Card{}}
	for id := range pokerCards {
		board.Cards[id] = pokerCards[id]
	}
	return board
}

func main() {
	game1 := NewPokerGame()
	game2 := NewPokerGame()
	fmt.Println(game1.Cards[1] == game2.Cards[1])
}
