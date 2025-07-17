package internal

import (
	"github.com/eiannone/keyboard"
	"log"
)

// NewGameLoop creates board and starts infinite loop with keyboard listen.
func NewGameLoop() {
	board := NewBoard2048(Size{4, 4}, nil)

	if err := board.Initialize(); err != nil {
		panic(err)
	}

	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			log.Fatal(err)
		}

		if err = board.Move(KeyToBoardEvent(key)); err != nil {
			log.Fatal(err)
		}

		// Handle match results.
		if board.IsWon() {
			log.Fatal("You won game!")
		}

		if board.IsLost() {
			log.Fatal("You lost.")
		}

	}
}
