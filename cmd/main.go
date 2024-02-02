package main

import (
	"fmt"
	"strings"
	"time"

	life "github.com/Noiidor/go-game-of-life/life"
)

func main() {
	field := []string{
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠ▮▮ᅠ▮ᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠ▮▮ᅠ▮ᅠ▮▮ᅠ▮▮▮ᅠ",
		"ᅠᅠ▮▮▮▮ᅠᅠ▮▮ᅠᅠᅠᅠᅠᅠ▮",
		"ᅠ▮ᅠᅠᅠᅠ▮ᅠᅠᅠ▮ᅠᅠᅠ▮▮ᅠ",
		"ᅠᅠ▮▮ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠ▮▮ᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠ▮ᅠ▮ᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠ▮ᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
		"ᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠᅠ",
	}

	for {
		fmt.Printf("\033[0;0H")
		fmt.Print("\r"+strings.Join(field, "\n"), "\n", strings.Repeat("-", len(field[0])))

		field = life.ProcessFieldIteration(field)

		time.Sleep(time.Millisecond * 50)

	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("\033[0;0H")

	// 	fmt.Print(strings.Join(field, "\n"))

	// 	time.Sleep(time.Second) // Simulating some work being done
	// }

	// fmt.Println("\nTask completed!")
}
