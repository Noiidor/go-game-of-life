package life

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func StartGameOfLife(iterations int, delayMs int) {
	// field := [][]bool{
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, true, true, false, false, false, false, false, false},
	// 	{false, false, false, false, true, false, true, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, true, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// 	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	// }

	field := RandomCellsInit(20, 20, 30)

	carriageReturneOnce := false
	if iterations < 1 { //eeehh..seems like not DRY at all
		iterCounter := 0
		for {

			fieldView := buildFieldString(field, iterCounter)

			if carriageReturneOnce { // hack to not let carriage to mess up terminal in first iteration
				height := len(field) + 3
				fmt.Printf("\033[%dA", height) // returns carriage up to field height(2 is field frame lines)

			} else {
				carriageReturneOnce = true
			}

			//fmt.Print(string('\x1b') + "[" + "s") // Saving cursor position for some reason dont work in powershell

			fmt.Print(fieldView)

			//fmt.Print(string('\x1b') + "[" + "u")

			field = processFieldIteration(field)

			iterCounter++

			time.Sleep(time.Millisecond * time.Duration(delayMs))

		}
	} else {
		for i := 0; i < iterations; i++ {
			fieldView := buildFieldString(field, i)

			if carriageReturneOnce {
				height := len(field) + 3
				fmt.Printf("\033[%dA", height)
			} else {
				carriageReturneOnce = true
			}

			fmt.Print(fieldView)

			field = processFieldIteration(field)

			time.Sleep(time.Millisecond * time.Duration(delayMs))
		}
	}
}

func processFieldIteration(field [][]bool) [][]bool {

	updatedField := copyFieldSlice(field)

	gridNeighboursCoords := [3]int{-1, 0, 1} // user to determine coords of neighbour cells

	for i, line := range field {
		for j, cell := range line {

			liveNeighboursCount := 0

			for _, coordChangeY := range gridNeighboursCoords {
				for _, coordChangeX := range gridNeighboursCoords {
					if coordChangeY == 0 && coordChangeX == 0 { // current cell
						continue
					}

					yIndex := loopedIndex(i+coordChangeY, len(field)) // make field looped at edges
					xIndex := loopedIndex(j+coordChangeX, len(line))

					isLive := field[yIndex][xIndex]

					if isLive {
						liveNeighboursCount++
					}
				}
			}

			if cell { // if cell is live
				if liveNeighboursCount > 3 || liveNeighboursCount < 2 {
					updatedField[i][j] = false
				}
			} else {
				if liveNeighboursCount == 3 {
					updatedField[i][j] = true
				}
			}
		}
	}

	return updatedField
}

func copyFieldSlice(field [][]bool) [][]bool {
	copiedField := make([][]bool, len(field))

	for i := range copiedField {
		copiedField[i] = make([]bool, len(field[i]))

		copy(copiedField[i], field[i])
	}

	return copiedField
}

func RandomCellsInit(fieldHeight int, fieldWidth int, livePercentage int) [][]bool {
	field := make([][]bool, fieldHeight)
	for i := 0; i < fieldHeight; i++ {
		field[i] = make([]bool, fieldWidth)
		for j := 0; j < fieldWidth; j++ {

			field[i][j] = rand.Intn(100) <= livePercentage
		}

	}
	return field
}

func buildFieldString(field [][]bool, iteration int) string {
	var builder strings.Builder

	liveCells := 0

	builder.WriteRune('╔')

	for i := 0; i < len(field[0]); i++ {
		builder.WriteString("══")
	}

	builder.WriteRune('╗')
	builder.WriteString("\n")

	for _, line := range field {
		builder.WriteRune('║')
		for _, cell := range line {
			if cell {
				builder.WriteString("██")
				liveCells++
			} else {
				builder.WriteString("  ")
			}
		}
		builder.WriteRune('║')
		builder.WriteString("\n")
	}

	builder.WriteRune('╚')

	for i := 0; i < len(field[0]); i++ {
		builder.WriteString("══")
	}

	builder.WriteRune('╝')
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("Iteration: %d Live cells: %d\n", iteration, liveCells))

	return builder.String()
}

func loopedIndex(index int, sliceLen int) int {
	return (index%sliceLen + sliceLen) % sliceLen
}

func ValidateField(field []string) bool {
	return true
}
