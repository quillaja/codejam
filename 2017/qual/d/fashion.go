// https://code.google.com/codejam/contest/dashboard?c=3264486#s=p3

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var t int
	fmt.Scanln(&t)
	t = 10

	var n, m, row, col int
	var kind string
	for i := 1; i <= t; i++ {
		fmt.Scanln(&n, &m)
		g := makeGrid(n)
		for j := 0; j < m; j++ {
			_, err := fmt.Scanln(&kind, &row, &col)
			if err != nil {
				panic(err)
			}
			g.set(row, col, kind[0])
		}

		fmt.Printf("Case #%d: %d %d (%dx%d, m=%d)\n", i, g.score(), -1, n, n, m)
		// solve(g)
	}
}

func makeGrid(n int) *grid {
	return &grid{
		data: bytes.Repeat([]byte{blank}, n*n),
		size: n,
	}
	// grid := make([][]string, 0)
	// for i := 0; i < n; i++ {
	// 	grid = append(grid, make([]string, n))
	// }
	// return grid
}

const (
	blank = '.'
	rook  = '+'
	bish  = 'x'
	best  = 'o'
)

type grid struct {
	data []byte
	size int
}

func (g *grid) get(r, c int) byte {
	return g.data[(c-1)+(r-1)*g.size]
}

func (g *grid) set(r, c int, kind byte) {
	g.data[(c-1)+(r-1)*g.size] = kind
}

func (g *grid) score() int {
	s := 0
	for _, kind := range g.data {
		switch kind {
		case rook, bish:
			s++
		case best:
			s += 2
		}
	}
	return s
}

func (g *grid) String() string {
	s := ""
	for r := 0; r < g.size; r++ {
		s += fmt.Sprintf("%s\n", g.data[r*g.size:(r+1)*g.size])
	}
	return s
}

func solve(g *grid) {
	fmt.Println(g, "\n")
}
