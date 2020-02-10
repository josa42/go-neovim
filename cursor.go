package neovim

type Cursor [2]int

func (c Cursor) X() int {
	return c[1]
}

func (c Cursor) Y() int {
	return c[0]
}

func (c *Cursor) SetX(x int) {
	c[1] = x
}

func (c *Cursor) SetY(y int) {
	c[0] = y
}

