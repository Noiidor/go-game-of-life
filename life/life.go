package life

const (
	deadCell = 'ᅠ'
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

func loopedIndex(index int, sliceLen int) int {
	return (index%sliceLen + sliceLen) % sliceLen
}

func ValidateField(field []string) bool {
	return true
}
