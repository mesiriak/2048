package internal

import (
	"fmt"
)

type Tiles [4][4]int

type Size struct {
	Width  int
	Height int
}

type Board interface {
	Step() int
	Size() Size
	Tiles() Tiles
	IsWon() bool
	IsLost() bool

	Initialize() error
	Move(event BoardEvent) error

	PrettyPrint(event BoardEvent)
}

type Board2048 struct {
	generator TilesGenerator
	tiles     Tiles
	size      Size
	step      int
}

func NewBoard2048(size Size, generator TilesGenerator) Board {
	return &Board2048{
		generator: generator,
		tiles:     Tiles{},
		size:      size,
		step:      1,
	}
}

func (b *Board2048) Step() int {
	return b.step
}

func (b *Board2048) Size() Size {
	return b.size
}

func (b *Board2048) Tiles() Tiles {
	return b.tiles
}

func (b *Board2048) IsWon() bool {
	for _, row := range b.tiles {
		for _, cell := range row {
			if cell == 2048 {
				return true
			}
		}
	}

	return false
}

func (b *Board2048) IsLost() bool {
	// if any cell value zero - not lose
	for _, row := range b.tiles {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}

	// if still any possible turns
	possibleTurns := []BoardEvent{Right, Down, Left, Up}

	for _, event := range possibleTurns {
		boardCopy := *b

		if err := boardCopy.Move(event); err != nil {
			panic(err)
		}

		if boardCopy.tiles != b.tiles {
			return false
		}
	}

	return true
}

// Initialize generates tiles for first step.
func (b *Board2048) Initialize() error {
	if b.generator == nil {
		b.generator = DefaultTilesGenerator
	}

	if err := b.generator(b.step, &b.tiles, b.size); err != nil {
		return err
	}

	b.PrettyPrint(Init)

	return nil
}

func (b *Board2048) PrettyPrint(event BoardEvent) {
	fmt.Println(fmt.Sprintf("EVENT - %s, STEP - %d", event, b.step))
	for _, row := range b.Tiles() {
		for _, cell := range row {
			fmt.Printf("%5d", cell)
		}
		fmt.Println()
	}
}

// Move rotates board in correct way for ease right-side merge, merges values, rotates board
// again, generates new tile on the board and increase step counter.
func (b *Board2048) Move(event BoardEvent) error {

	rotationCount, err := eventToRotationCount(event)
	if err != nil {
		return err
	}

	invertedTiles := collapseToRight(rotateTilesToRight(b.tiles, b.size, rotationCount), b.size)

	invertedRotationCount, err := eventToInvertedRotationCount(event)
	if err != nil {
		return err
	}

	tiles := rotateTilesToRight(invertedTiles, b.size, invertedRotationCount)

	// if step did nothing - do nothing
	if tiles != b.tiles {
		b.tiles = tiles
		b.step++

		if err = b.generator(b.step, &b.tiles, b.size); err != nil {
			return err
		}

		b.PrettyPrint(event)
	}

	return nil
}
