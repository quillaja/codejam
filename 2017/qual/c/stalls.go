// https://code.google.com/codejam/contest/dashboard?c=3264486#s=p2
// first impression is that this is a binary tree type problem
// but with N up to 10^18, this can't store "state" (it's an exabit
// of memory at a minimum)

package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanln(&t)
	// t = 15
	var n, k uint64
	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		h, l := solve2(n, k)
		fmt.Printf("Case #%d: %d %d\n", i+1, h, l)
		// fmt.Printf("Case #%d:(n=%d k=%d) %d %d\n", i+1, n, k, h, l)
	}
}

// seemed maybe closer.. but still wrong.
func solve2(n, k uint64) (h, l uint64) {
	depth := level(k)
	return n / pow2(depth+1), (n - 1) / pow2(depth+1)
}

// i don't know why this isn't working. it seems to work for all the
// "explained" test cases (1-5) and some I made up.
//
// ... except it didnt: case 3 (n=6, k=2) was wrong. I'm getting (1,0), and the
// answer is (1,1). On paper I'm getting this of course, but my little "formula"
// isn't working for SOME cases when n is even (seems to work when n=1000...)
func solve(n, k uint64) (h, l uint64) {
	// level := int(math.Ceil(math.Log2(float64(k+1)) - 1))
	// height := level(k)
	// // fmt.Println(level)
	// h, l = n/2, (n-1)/2
	// if height == 0 {
	// 	return
	// }

	// // go right
	// solve(max(h, l), k-above(height))

	// // go left
	// solve(min(h, l), k-above(height))

	todepth := level(k)
	curdepth := uint64(0)
	visited := uint64(0)
	for curdepth <= todepth {
		if n%2 == 0 && n > 0 {
			h = n / 2
			l = (n - 1) / 2
		} else {
			h = n / 2
			l = n / 2
		}
		// fmt.Println(h, l)

		visited = above(curdepth)
		curdepth++

		// TODO: I don't have any idea how to figure out to "go right"
		// or "go left". I think I'd have the problem solved if I could
		// figure that out.
		if (k-visited)%2 == 1 { // go right
			// fmt.Println("went right")
			n = max(h, l)
			// k = k - 1
		} else { // go left
			// fmt.Println("went left")
			n = min(h, l)
			// k = k - 2
		}
		// fmt.Println(h, l, n, visited)
	}

	return
}

// level of kth (1,2,...) stall taken
func level(k uint64) uint64 {
	return uint64(math.Floor(math.Log2(float64(k))))
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func pow2(x uint64) uint64 {
	return 1 << x
}

func above(h uint64) uint64 {
	return pow2(h+1) - 1
}
