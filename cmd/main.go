package main

import (
	"fmt"
	"time"

	life "github.com/Noiidor/go-game-of-life/life"
)

func main() {
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

	for {
		//fmt.Printf("\033[0;0H")
		fieldView := life.BuildFieldString(field)

		fmt.Print(fieldView)

		field = life.ProcessFieldIterationNew(field)

		time.Sleep(time.Millisecond * 100)

	}
}
