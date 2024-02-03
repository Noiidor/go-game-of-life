package life

import "strings"

const (
	deadCell = '—'
	liveCell = '▮'
)

func ProcessFieldIteration(field []string) []string {
	var updatedField []string

	for i := 0; i < len(field); i++ {
		rowsCount := len(field)

		currentLine := []rune(field[i])
		previousLine := []rune(field[loopedIndex(i-1, rowsCount)])
		nextLine := []rune(field[loopedIndex(i+1, rowsCount)])

		updatedLine := make([]rune, len(currentLine))
		copy(updatedLine, currentLine)

		for j := 0; j < len(currentLine); j++ {

			lineLength := len(currentLine)

			isCellAlive := currentLine[j] == liveCell
			liveNeighboursCount := 0

			leftCell := currentLine[loopedIndex(j-1, lineLength)]
			if leftCell == liveCell {
				liveNeighboursCount++
			}

			rightCell := currentLine[loopedIndex(j+1, lineLength)]
			if rightCell == liveCell {
				liveNeighboursCount++
			}

			leftTopCell := previousLine[loopedIndex(j-1, lineLength)]
			if leftTopCell == liveCell {
				liveNeighboursCount++
			}

			centralTopCell := previousLine[loopedIndex(j, lineLength)]
			if centralTopCell == liveCell {
				liveNeighboursCount++
			}

			rightTopCell := previousLine[loopedIndex(j+1, lineLength)]
			if rightTopCell == liveCell {
				liveNeighboursCount++
			}

			leftBottomCell := nextLine[loopedIndex(j-1, lineLength)]
			if leftBottomCell == liveCell {
				liveNeighboursCount++
			}

			centralBottomCell := nextLine[loopedIndex(j, lineLength)]
			if centralBottomCell == liveCell {
				liveNeighboursCount++
			}

			rightBottomCell := nextLine[loopedIndex(j+1, lineLength)]
			if rightBottomCell == liveCell {
				liveNeighboursCount++
			}

			if isCellAlive {
				if liveNeighboursCount > 3 || liveNeighboursCount < 2 {
					updatedLine[j] = deadCell
				}
			} else {
				if liveNeighboursCount == 3 {
					updatedLine[j] = liveCell
				}
			}

		}

		updatedField = append(updatedField, string(updatedLine))

	}
	return updatedField
}

func ProcessFieldIterationNew(field [][]bool) [][]bool {

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

func BuildFieldString(field [][]bool) string {
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
