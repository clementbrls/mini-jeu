package main

type DiscColor int

const (
	Empty DiscColor = iota
	Black
	White
)

type Cell struct {
	color DiscColor
}

func (c Cell) isEmpty() bool {
	return c.color == Empty
}

func (c Cell) isBlack() bool {
	return c.color == Black
}

func (c Cell) isWhite() bool {
	return c.color == White
}
