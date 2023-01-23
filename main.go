package main

import (
	"fmt"
	"sarge424/sudokusolver/board"
	"time"
)

func main() {
	b := board.New()
	fmt.Println("Initializing...")

	fmt.Print("Use example? (y/n):")
	text := ""
	fmt.Scanln(&text)

	if text[0] == 'y' {
		b.SetRow(0, "150000002")
		b.SetRow(1, "040800150")
		b.SetRow(2, "000050007")

		b.SetRow(3, "300060540")
		b.SetRow(4, "006200000")
		b.SetRow(5, "000000070")

		b.SetRow(6, "080009000")
		b.SetRow(7, "000000001")
		b.SetRow(8, "400020360")

	} else {
		for i := 0; i < 9; i++ {
			fmt.Print("row", i+1, ":")
			fmt.Scanln(&text)
			b.SetRow(i, text)
		}
	}

	b.Collapse()

	b.Print(false)

	start := time.Now()
	fmt.Print("\nSolving...")

	b.Solve()

	elapsed := time.Since(start)

	fmt.Println("Done.\nSolved in", elapsed)
	b.Print(false)

	fmt.Print("\nPress enter to exit.")
	fmt.Scanln(&text)
}
