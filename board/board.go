package board

import (
	"fmt"
	"sarge424/sudokusolver/cell"
)

type Board struct {
	data    [9][9]*cell.Cell
	unknown int
	solved  bool
}

func New() Board {
	b := Board{}
	b.unknown = 81
	for i, row := range b.data {
		for j := range row {
			b.data[i][j] = cell.New()
		}
	}

	return b
}

func (b *Board) SetRow(r int, s string) {
	for j, c := range s {
		x := int(c - '0')
		b.data[r][j].SetValue(x)
		b.data[r][j].Lock()
	}

	b.unknown = b.Blanks()
}

func (b Board) getRowValues(row int, col int) [9]bool {
	var a [9]bool

	for i, c := range b.data[row] {
		if i == col { //skip current cell
			continue
		}
		if c.GetValue() != -1 {
			a[c.GetValue()] = true
		}
	}

	return a
}

func (b Board) getColValues(row int, col int) [9]bool {
	var a [9]bool

	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		c := b.data[i][col]
		if c.GetValue() != -1 {
			a[c.GetValue()] = true
		}
	}

	return a
}

func (b Board) getSquareValues(row int, col int) [9]bool {
	var a [9]bool
	rowIndex := int(row/3) * 3
	colIndex := int(col/3) * 3

	for i := rowIndex; i < rowIndex+3; i++ {
		for j := colIndex; j < colIndex+3; j++ {
			//skip current cell
			if i == row && j == col {
				continue
			}
			val := b.data[i][j].GetValue()
			if val >= 0 {
				a[val] = true
			}
		}
	}

	return a
}

func invert(a [9]bool) [9]bool {
	for i := range a {
		a[i] = !a[i]
	}

	return a
}

func intersect(doms ...[9]bool) [9]bool {
	var a [9]bool
	for i := 0; i < 9; i++ {
		a[i] = true
	}

	for _, dom := range doms {
		for i, val := range dom {
			a[i] = a[i] && val
		}
	}

	return a
}

func (b *Board) Collapse() {
	for i, row := range b.data {
		for j := range row {
			rowData := invert(b.getRowValues(i, j))
			colData := invert(b.getColValues(i, j))
			squareData := invert(b.getSquareValues(i, j))
			b.data[i][j].SetDomain(intersect(rowData, colData, squareData))
		}
	}
}

func (b Board) IsValid() bool {
	var rows [9][10]int
	var cols [9][10]int
	var squares [3][3][10]int

	for i, row := range b.data {
		for j, c := range row {
			val := c.GetValue() + 1
			ri := int(i / 3)
			ci := int(j / 3)
			rows[i][val]++
			cols[j][val]++
			squares[ri][ci][val]++
			if rows[i][val] > 1 || cols[j][val] > 1 || squares[ri][ci][val] > 1 {
				if val > 0 {
					return false
				}
			}

		}
	}

	return true
}

func (b *Board) RecSolve(digit int) {
	x := digit + 1
	if digit >= b.unknown {
		if b.IsValid() {
			b.solved = true
			return
		}
	}
	for _, row := range b.data {
		for _, c := range row {
			if x == 1 && !c.Locked() {
				b.Collapse()
				//go through its domain and check each
				_, dom := c.GetAll()
				for i, d := range dom {
					if d {
						c.SetValue(i + 1)
						//fmt.Println("on", p+1, q+1, i+1)
						//fmt.Println(spaces(digit-1), "substituting", i+1, "in", p+1, q+1)
						b.RecSolve(digit + 1)

						//b.Print(true)

						if b.solved {
							return
						}
					}
					c.SetValue(0)
					b.Collapse()
				}

				return
			} else if !c.Locked() {
				x--
			}
		}
	}
}

func (b *Board) Solve() {
	if b.IsValid() {
		b.RecSolve(0)
	} else {
		fmt.Println("The given puzzle is invalid.")
	}
}

func (b Board) Blanks() int {
	blanks := 0
	for _, row := range b.data {
		for _, c := range row {
			if c.GetValue() < 0 {
				blanks++
			}
		}
	}

	return blanks
}

func (b Board) Print(debug bool) {
	var combs uint64 = 1
	if !debug {
		fmt.Println("┌───────┬───────┬───────┐")
	}
	for i, row := range b.data {

		if !debug {
			fmt.Print("│ ")
		}

		for j, c := range row {
			if debug {
				c.Debug()
			} else {
				c.Print()
			}

			if j%3 == 2 && !debug {
				fmt.Print("│ ")
			}

			combs *= uint64(c.Length())
		}
		fmt.Println()
		if i%3 == 2 && i < 8 && !debug {
			fmt.Println("├───────┼───────┼───────┤")
		}
	}
	if !debug {
		fmt.Println("└───────┴───────┴───────┘")
	}
	fmt.Println("possiblities  :", combs)
	fmt.Println("Valid?        :", b.IsValid())
	fmt.Println("Solved?       :", b.solved)
	fmt.Println("unknown       :", b.Blanks())
}
