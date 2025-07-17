package internal

import "math/rand"

type TilesGenerator func(step int, tiles *Tiles, size Size) error

// DefaultTilesGenerator make 3 tiles in default positions on start
// and generate tile in random cell in other turns.
func DefaultTilesGenerator(step int, tiles *Tiles, size Size) error {
	possibleStartValues := []int{2, 4}

	if step == 1 {
		for i := 0; i < 3; i++ {
			row, col := rand.Intn(size.Height), rand.Intn(size.Width)

			if tile := tiles[row][col]; tile == 0 {
				tiles[row][col] = possibleStartValues[rand.Intn(len(possibleStartValues))]
			}
		}

		return nil
	}

	var freeCells [][2]int

	possibleStartValues = append(possibleStartValues, 8)

	for row := 0; row < size.Height; row++ {
		for col := 0; col < size.Width; col++ {
			if tile := tiles[row][col]; tile == 0 {
				freeCells = append(freeCells, [2]int{row, col})
			}
		}
	}

	if len(freeCells) > 0 {
		rand.Shuffle(len(freeCells), func(i, j int) {
			freeCells[i], freeCells[j] = freeCells[j], freeCells[i]
		})

		tiles[freeCells[0][0]][freeCells[0][1]] = possibleStartValues[rand.Intn(len(possibleStartValues))]
	}

	return nil
}
