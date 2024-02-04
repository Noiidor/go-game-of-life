package life

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func StartGameOfLife(iterations int, delayMs int, fieldHeight int, fieldWidth int, startLiveCellPercent int) {
	field := randomCellsInit(fieldHeight, fieldWidth, startLiveCellPercent)

	infiniteLoop := iterations < 1
	liveCells := 0

	for i := 0; infiniteLoop || i < iterations; i++ {
		printFieldView(field, i, liveCells)

		field, liveCells = processFieldIteration(field)

		time.Sleep(time.Millisecond * time.Duration(delayMs))
	}
}

func printFieldView(field [][]bool, iteration int, liveCells int) {
	fieldView := buildFieldString(field, iteration)

	if iteration != 0 { // hack to not let carriage to mess up terminal in first iteration
		height := len(field) + 3
		fmt.Printf("\033[%dA", height)
	}

	fmt.Print(fieldView)
	fmt.Printf("Iteration: %d Live cells: %v     \n", iteration, liveCells)
} //                                          ^ workaround to override printed digits from last iteration with whitespaces

func processFieldIteration(field [][]bool) ([][]bool, int) {

	updatedField := copyFieldSlice(field)

	gridNeighboursCoords := [3]int{-1, 0, 1} // user to determine coords of neighbour cells

	currentLiveCells := 0

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
				} else {
					currentLiveCells++
				}
			} else {
				if liveNeighboursCount == 3 {
					updatedField[i][j] = true
					currentLiveCells++
				}
			}
		}
	}

	return updatedField, currentLiveCells
}

func copyFieldSlice(field [][]bool) [][]bool {
	copiedField := make([][]bool, len(field))

	for i := range copiedField {
		copiedField[i] = make([]bool, len(field[i]))

		copy(copiedField[i], field[i])
	}

	return copiedField
}

func randomCellsInit(fieldHeight int, fieldWidth int, livePercentage int) [][]bool {
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

	return builder.String()
}

func loopedIndex(index int, sliceLen int) int {
	return (index%sliceLen + sliceLen) % sliceLen
}
