package internal

// rotateTilesToRight returns a new Tiles array that is the result of rotating
// the given tiles 90 degrees to the right (clockwise) rotationCount times.
func rotateTilesToRight(tiles Tiles, size Size, rotationCount int) Tiles {
	if rotationCount == 0 {
		return tiles
	}

	var rotatedTiles Tiles

	for i := 0; i < rotationCount; i++ {

		for col := 0; col < size.Height; col++ {
			var rotatedRow [4]int

			for row := size.Width - 1; row >= 0; row-- {
				rotatedRow[size.Width-row-1] = tiles[row][col]
			}

			rotatedTiles[col] = rotatedRow
		}
		tiles, rotatedTiles = rotatedTiles, Tiles{}
	}

	return tiles
}

func collapseToRight(tiles Tiles, size Size) Tiles {
	for col := size.Height - 2; col >= 0; col-- {
		for row := 0; row < size.Width; row++ {
			if tiles[row][col] == 0 {
				continue
			}

			if tiles[row][col] != 0 {
				mergeToRight(row, col, &tiles, size)
			}
		}
	}

	return tiles
}

func mergeToRight(row, col int, tiles *Tiles, size Size) {
	if col == size.Width-1 {
		return
	}

	// till next equal 0 - move to right
	if tiles[row][col+1] == 0 {
		tiles[row][col+1], tiles[row][col] = tiles[row][col], 0
		mergeToRight(row, col+1, tiles, size)
	}

	// if next not equal 0 and next same as current - merge
	if tiles[row][col+1] != 0 && tiles[row][col+1] == tiles[row][col] {
		tiles[row][col+1], tiles[row][col] = tiles[row][col]+tiles[row][col+1], 0
		return
	}

	// if next not equal and next not same as current - return
	if tiles[row][col+1] != 0 && tiles[row][col+1] != tiles[row][col] {
		return
	}

	mergeToRight(row, col+1, tiles, size)
}
