package main

import (
	"github.com/eiannone/keyboard"
	"github.com/mesiriak/2048/internal"
)

// Initialize keyboard CLI, start game loop.
func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() { _ = keyboard.Close() }()

	internal.NewGameLoop()

}
