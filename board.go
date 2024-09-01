package main

import "fmt"

type Board struct {
	cells [8][8]Cell
}

type coordinates struct {
	x int
	y int
}

func (b *Board) Initialize() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.cells[i][j] = Cell{color: Empty}
		}
	}
	b.cells[3][3] = Cell{color: White}
	b.cells[4][4] = Cell{color: White}
	b.cells[3][4] = Cell{color: Black}
	b.cells[4][3] = Cell{color: Black}
}

func (b Board) String() string {
	str := "---------------------------------\n"
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if b.cells[i][j].isEmpty() {
				str += "|   "
			} else if b.cells[i][j].isBlack() {
				str += "| X "
			} else if b.cells[i][j].isWhite() {
				str += "| O "
			}
		}
		str += "|\n"
		str += "---------------------------------\n"
	}
	return str
}

func (b *Board) Play(x, y int, color DiscColor) {
	validMoves := b.validMoves(color)
	fmt.Printf("Valid moves: %v\n", validMoves)
	coord := coordinates{x: x, y: y}
	if b.isValidMove(coord, validMoves) {
		b.cells[x][y] = Cell{color: color}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				} else if b.testDirection(x, y, dx, dy, color) {
					b.flip(x, y, dx, dy, color)
				}
			}
		}
	} else {
		fmt.Errorf("Invalid move")
	}
}

func (b Board) testDirection(x int, y int, dx int, dy int, color DiscColor) bool {
	oppositeColor := Black
	if color == Black {
		oppositeColor = White
	}

	if x+dx < 0 || x+dx >= len(b.cells) || y+dy < 0 || y+dy >= len(b.cells[0]) || (dy == 0 && dx == 0) {
		return false
	} else if b.cells[x+dx][y+dy].color != oppositeColor {
		return false
	}

	for i := 2; !(x+i*dx < 0 || x+i*dx >= len(b.cells) || y+i*dy < 0 || y+i*dy >= len(b.cells[0])); i++ {
		if b.cells[x+i*dx][y+i*dy].color == color {
			return true
		} else if b.cells[x+i*dx][y+i*dy].color == Empty {
			return false
		}
	}
	return false
}

func (b *Board) flip(x int, y int, dx int, dy int, color DiscColor) {
	oppositeColor := Black
	if color == Black {
		oppositeColor = White
	}

	for i := 1; !(x+i*dx < 0 || x+i*dx >= len(b.cells) || y+i*dy < 0 || y+i*dy >= len(b.cells[0])); i++ {
		if b.cells[x+i*dx][y+i*dy].color == oppositeColor {
			b.cells[x+i*dx][y+i*dy] = Cell{color: color}
		} else if b.cells[x+i*dx][y+i*dy].color == color {
			break
		} else {
			break
		}
	}
}

func (b Board) validMoves(color DiscColor) []coordinates {
	validMoves := []coordinates{}
	for x := 0; x < len(b.cells); x++ {
		for y := 0; y < len(b.cells[0]); y++ {
			if b.cells[x][y].isEmpty() {
			out:
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						} else if b.testDirection(x, y, dx, dy, color) {
							validMoves = append(validMoves, coordinates{x: x, y: y})
							break out
						}
					}
				}
			}
		}
	}
	return validMoves
}

func (b Board) isValidMove(coord coordinates, validMoves []coordinates) bool {
	for _, move := range validMoves {
		if coord == move {
			return true
		}
	}
	return false
}
