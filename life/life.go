package life

import (
	"strings"
	"fmt"
	"time"
)

func StartGameOfLife(iterations int, delayMs int){
	field := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}

	if iterations < 1 {
		for {
			fieldView := buildFieldString(field)
			//fmt.Printf("\033[0;0H")
			fmt.Print(fieldView)

			field = processFieldIteration(field)

			time.Sleep(time.Millisecond * time.Duration(delayMs))
		}
	}else{
		for i := 0; i < iterations; i++ {
			fieldView := buildFieldString(field)
			//fmt.Printf("\033[0;0H")
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

func buildFieldString(field [][]bool) string {
	var builder strings.Builder

	for _, line := range field {
		for _, cell := range line {
			if cell {
				builder.WriteString("██")
			} else {
				builder.WriteString("  ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func loopedIndex(index int, sliceLen int) int {
	return (index%sliceLen + sliceLen) % sliceLen
}

func ValidateField(field []string) bool {
	return true
}
