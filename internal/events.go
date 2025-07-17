package internal

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type BoardEvent string

const (
	Left  BoardEvent = "left"
	Right BoardEvent = "right"
	Up    BoardEvent = "up"
	Down  BoardEvent = "down"
	Other BoardEvent = "other"
	Init  BoardEvent = "init"
)

func KeyToBoardEvent(key keyboard.Key) BoardEvent {
	switch key {
	case keyboard.KeyArrowUp:
		return Up
	case keyboard.KeyArrowDown:
		return Down
	case keyboard.KeyArrowLeft:
		return Left
	case keyboard.KeyArrowRight:
		return Right
	default:
		fmt.Println("Please, use only arrows.")
		return Other
	}
}

func eventToRotationCount(event BoardEvent) (int, error) {
	switch event {
	case Right:
		return 0, nil
	case Up:
		return 1, nil
	case Left:
		return 2, nil
	case Down:
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid event: %s", event)
	}
}

func eventToInvertedRotationCount(event BoardEvent) (int, error) {
	switch event {
	case Right:
		return 0, nil
	case Down:
		return 1, nil
	case Left:
		return 2, nil
	case Up:
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid event: %s", event)
	}
}
