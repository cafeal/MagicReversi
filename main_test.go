package main

import (
	"log"
	"testing"
)

type dammyMiddleware struct {
	// a number of current turn
	t int
	// test input data set
	r [][2]int
}

func (m *dammyMiddleware) GetInput() (x int, y int) {
	n := m.r[m.t]
	m.t++

	x, y = n[0], n[1]

	return
}

func (m *dammyMiddleware) Flip(x int, y int) {
	// do nothing
}

func TestMain(t *testing.T) {
	m := &dammyMiddleware{
		t: 0,
		// test data
		r: [][2]int{
			[2]int{3, 4},
			[2]int{3, 3},
			[2]int{4, 3},
			[2]int{5, 3},
			[2]int{4, 2},
			[2]int{2, 4},
			[2]int{6, 4},
			[2]int{6, 3},
			[2]int{4, 6},
			[2]int{5, 6},
			[2]int{6, 2},
			[2]int{3, 6},
			[2]int{3, 5},
			[2]int{2, 5},
			[2]int{1, 4},
			[2]int{1, 5},
			[2]int{5, 2},
			[2]int{4, 1},
			[2]int{1, 6},
			[2]int{2, 6},
			[2]int{2, 3},
			[2]int{3, 2},
			[2]int{6, 5},
			[2]int{1, 2},
			[2]int{6, 7},
			[2]int{6, 1},
			[2]int{5, 1},
			[2]int{5, 7},
			[2]int{1, 3},
			[2]int{1, 7},
			[2]int{6, 6},
			[2]int{7, 5},
			[2]int{3, 7},
			[2]int{6, 8},
			[2]int{5, 8},
			[2]int{7, 3},
			[2]int{7, 4},
			[2]int{7, 6},
			[2]int{3, 1},
			[2]int{8, 4},
			[2]int{7, 1},
			[2]int{2, 2},
			[2]int{7, 8},
			[2]int{3, 8},
			[2]int{8, 3},
			[2]int{8, 2},
			[2]int{8, 6},
			[2]int{8, 5},
			[2]int{8, 1},
			[2]int{7, 2},
			[2]int{2, 7},
			[2]int{4, 8},
			[2]int{2, 8},
			[2]int{4, 7},
			[2]int{1, 8},
			[2]int{8, 7},
			[2]int{1, 1},
			[2]int{7, 7},
			[2]int{2, 1},
			[2]int{8, 8},
		},
	}

	g := newGame(m)

	err := g.start()

	if err != nil {
		log.Fatal(err)
	}
}