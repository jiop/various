package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const MAX_Y = 15
const MAX_X = 100
const NB_CELLS = 50

type Board struct {
	grid    [2][MAX_X][MAX_Y]bool
	current int
}

type Cell struct {
	x int
	y int
}

func newBoard() *Board {
	return &Board{}
}

func (b *Board) seed() {
	for i := 0; i < NB_CELLS; i++ {
		b.grid[b.current][rand.Intn(MAX_X)][rand.Intn(MAX_Y)] = true
	}
}

func (b *Board) print() {
	for j := 0; j < MAX_Y; j++ {
		for i := 0; i < MAX_X; i++ {
			if b.grid[b.current][i][j] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) getAliveNeighboursCount(c Cell) int {
	alives := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if b.getCellState(Cell{x: c.x + i, y: c.y + j}) {
				alives++
			}
		}
	}
	return alives
}

func (b *Board) getCellState(c Cell) bool {
	if c.x < 0 || c.y < 0 || c.x >= MAX_X || c.y >= MAX_Y {
		return false
	}
	return b.grid[b.current][c.x][c.y]
}

func (b *Board) nextState(c Cell) bool {
	alives := b.getAliveNeighboursCount(c)
	if alives < 2 {
		return false // underpopulation
	}
	if alives > 3 {
		return false // overpopulation
	}
	if alives == 3 && !b.getCellState(c) {
		return true // reproduction
	}
	return true // survives
}

func (b *Board) step() {
	for j := 0; j < MAX_Y; j++ {
		for i := 0; i < MAX_X; i++ {
			b.grid[(b.current+1)%2][i][j] = b.nextState(Cell{x: i, y: j})
		}
	}
	b.current = (b.current + 1) % 2
}

func main() {
	rand.Seed(time.Now().UnixNano())
	b := newBoard()
	b.seed()

	for {
		select {
		case <-time.After(200 * time.Millisecond):
			b.step()
			fmt.Println(strings.Repeat("#", MAX_Y))
			b.print()
		}
	}
}
