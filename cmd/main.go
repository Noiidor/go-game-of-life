package main

import (
	"flag"

	life "github.com/Noiidor/go-game-of-life/life"
)

func main() {
	numOfIterations := flag.Int("iter", 0, "Number of iterations to run. 0 for infinite loop.")

	delay := flag.Int("delay", 100, "Delay between iterations(ms). Default is 100.")

	fieldHeight := flag.Int("height", 30, `Height of game field in cells. Default is 30. 
					In some terminals height or width that exceeds terminal borders will break game field display.`)

	fieldWidth := flag.Int("width", 30, "Width of game field in cells. Default is 30.")

	livePercentage := flag.Int("start-live", 35, "Percent of live cells in first iteration. Default is 35.")

	flag.Parse()

	life.StartGameOfLife(*numOfIterations, *delay, *fieldHeight, *fieldWidth, *livePercentage)
}
