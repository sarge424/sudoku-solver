package cell

import "fmt"

type Cell struct {
	value  int
	domain [9]bool
	locked bool
}

func New() *Cell {
	l := new(Cell)
	l.value = -1
	for i := range l.domain {
		l.domain[i] = true
	}
	return l
}

func (c *Cell) SetValue(num int) {
	if num > 0 && num < 10 {
		c.value = num - 1
	} else {
		c.value = -1
	}
}

func (c *Cell) Lock() {
	if c.value < 0 {
		return
	}
	c.locked = true
	for i := 0; i < 9; i++ {
		c.domain[i] = false
	}

	c.domain[c.value] = true
}

func (c *Cell) SetDomain(dom [9]bool) {
	if !c.locked {
		c.domain = dom
	}
}

func (c Cell) Locked() bool {
	return c.locked
}

func (c Cell) Length() int {
	if c.locked {
		return 1
	}
	a := 0
	for _, val := range c.domain {
		if val {
			a++
		}
	}

	return a
}

func (c Cell) GetAll() (int, [9]bool) {
	return c.value, c.domain
}

func (c Cell) GetValue() int {
	return c.value
}

func (c *Cell) Print() {
	if c.value == -1 {
		fmt.Print("  ")
	} else {
		fmt.Print(c.value+1, " ")
	}
}

func (c *Cell) Debug() {
	if c.locked {
		fmt.Print(c.value+1, " ")
		return
	}

	if c.value > -1 {
		fmt.Print(c.value + 1)
	}
	fmt.Print("[")
	for i, val := range c.domain {
		if val {
			fmt.Print(i + 1)
		}
	}
	fmt.Print("] ")
}
