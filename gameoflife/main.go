package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"strings"
	"time"
)

const maxY = 15
const maxX = 50
const nbCells = 10

type board struct {
	grid        [2][maxX][maxY]bool
	currentGrid int
	printer     printer
}

type printer interface {
	print(*board)
}

type termPrinter struct{}

func (termPrinter) print(b *board) {
	const clrR = "\x1b[31;1m"
	const clrG = "\x1b[32;1m"
	const clrN = "\x1b[0m"
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			if b.current()[i][j] {
				fmt.Printf("%sx%s", clrR, clrN)
			} else {
				fmt.Printf("%so%s", clrG, clrN)
			}
		}
		fmt.Println()
	}
}

type floydSteinbergPrinter struct{}

func (floydSteinbergPrinter) print(b *board) {
	im := image.NewGray(image.Rectangle{Max: image.Point{X: maxX, Y: maxY}})
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if b.getCellState(i, j) {
				im.SetGray(i, j, color.Gray{Y: 0})
			} else {
				im.SetGray(i, j, color.Gray{Y: 35})
			}
		}
	}
	pi := image.NewPaletted(im.Bounds(), []color.Color{
		color.Gray{Y: 255},
		color.Gray{Y: 160},
		color.Gray{Y: 70},
		color.Gray{Y: 35},
		color.Gray{Y: 0},
	})
	draw.FloydSteinberg.Draw(pi, im.Bounds(), im, image.ZP)
	shade := []string{" ", "░", "▒", "▓", "█"}
	for i, p := range pi.Pix {
		fmt.Print(shade[p])
		if (i+1)%maxX == 0 {
			fmt.Print("\n")
		}
	}
}

func newBoard(p printer) *board {
	return &board{printer: p}
}

func (b *board) current() [maxX][maxY]bool {
	return b.grid[b.currentGrid]
}

func (b *board) seed() {
	for i := 0; i < nbCells; i++ {
		b.grid[b.currentGrid][rand.Intn(maxX)][rand.Intn(maxY)] = true
	}
}

func (b *board) print() {
	b.printer.print(b)
}

func (b *board) getAliveNeighboursCount(x int, y int) int {
	alives := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if b.getCellState(x+i, y+j) {
				alives++
			}
		}
	}
	return alives
}

func (b *board) getCellState(x int, y int) bool {
	if x < 0 || y < 0 || x >= maxX || y >= maxY {
		return false
	}
	return b.current()[x][y]
}

func (b *board) nextState(x int, y int) bool {
	alives := b.getAliveNeighboursCount(x, y)
	if alives < 2 || alives > 3 {
		return false // underpopulation or overpopulation
	}
	if alives == 3 && !b.getCellState(x, y) {
		return true // reproduction
	}
	return true // survives
}

func (b *board) tick() {
	next := (b.currentGrid + 1) % 2
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			b.grid[next][i][j] = b.nextState(i, j)
		}
	}
	b.currentGrid = next
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// b := newBoard(&termPrinter{})
	b := newBoard(&floydSteinbergPrinter{})
	b.seed()

	for {
		select {
		case <-time.After(200 * time.Millisecond):
			b.tick()
			fmt.Println(strings.Repeat("#", maxX))
			b.print()
		}
	}
}
