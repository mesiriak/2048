package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mesiriak/2048/internal"
	"image/color"
	"log"
)

func pickTileColor(value int) color.NRGBA {
	switch value {
	case 0:
		return color.NRGBA{R: 205, G: 193, B: 180, A: 255}
	case 2:
		return color.NRGBA{R: 238, G: 228, B: 218, A: 255}
	case 4:
		return color.NRGBA{R: 237, G: 224, B: 200, A: 255}
	case 8:
		return color.NRGBA{R: 242, G: 177, B: 121, A: 255}
	case 16:
		return color.NRGBA{R: 245, G: 149, B: 99, A: 255}
	case 32:
		return color.NRGBA{R: 246, G: 124, B: 95, A: 255}
	case 64:
		return color.NRGBA{R: 246, G: 94, B: 59, A: 255}
	case 128:
		return color.NRGBA{R: 237, G: 207, B: 114, A: 255}
	case 256:
		return color.NRGBA{R: 237, G: 204, B: 97, A: 255}
	case 512:
		return color.NRGBA{R: 237, G: 200, B: 80, A: 255}
	case 1024:
		return color.NRGBA{R: 237, G: 197, B: 63, A: 255}
	case 2048:
		return color.NRGBA{R: 237, G: 194, B: 46, A: 255}
	default:
		return color.NRGBA{R: 60, G: 58, B: 50, A: 255}
	}
}

func main() {
	myWindow := app.New().NewWindow("2048 Game")

	boardSize := internal.Size{Width: 4, Height: 4}
	board := internal.NewBoard2048(boardSize, nil)

	if err := board.Initialize(); err != nil {
		log.Fatal(err)
	}

	var grid *fyne.Container

	updateGrid := func() {
		grid.Objects = nil
		for i := 0; i < boardSize.Height; i++ {
			for j := 0; j < boardSize.Width; j++ {
				value := board.Tiles()[i][j]

				rect := canvas.NewRectangle(pickTileColor(value))
				rect.SetMinSize(fyne.NewSize(60, 60))

				label := canvas.NewText(fmt.Sprintf("%d", value), color.Black)
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignCenter
				label.TextSize = 24

				cell := container.NewMax(rect, container.NewCenter(label))

				grid.Add(cell)
			}
		}
		grid.Refresh()
	}

	checkGameState := func() {
		if board.IsWon() {
			dialog.ShowInformation("You Win!", "Congratulations! You won the game!", myWindow)
		} else if board.IsLost() {
			dialog.ShowInformation("Game Over", "The game is over. You lost.", myWindow)
		}
	}

	moveAndUpdate := func(direction internal.BoardEvent) {
		err := board.Move(direction)
		if err != nil {
			log.Println("Move error:", err)
		}
		updateGrid()
		checkGameState()
	}

	grid = container.NewGridWithColumns(4)
	updateGrid()

	// Кнопки керування
	upBtn := widget.NewButton("Up", func() {
		moveAndUpdate(internal.Up)
	})

	downBtn := widget.NewButton("Down", func() {
		moveAndUpdate(internal.Down)
	})

	leftBtn := widget.NewButton("Left", func() {
		moveAndUpdate(internal.Left)
	})

	rightBtn := widget.NewButton("Right", func() {
		moveAndUpdate(internal.Right)
	})

	controlButtons := container.NewGridWithColumns(4, upBtn, downBtn, leftBtn, rightBtn)

	content := container.NewVBox(grid, controlButtons)
	myWindow.SetContent(content)

	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		switch keyEvent.Name {
		case fyne.KeyUp:
			moveAndUpdate(internal.Up)
		case fyne.KeyDown:
			moveAndUpdate(internal.Down)
		case fyne.KeyLeft:
			moveAndUpdate(internal.Left)
		case fyne.KeyRight:
			moveAndUpdate(internal.Right)
		}
	})

	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}
