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
	var n, k int
	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		h, l := solve3(n, k)
		fmt.Printf("Case #%d: %d %d\n", i+1, h, l)
		// fmt.Printf("Case #%d:(n=%d k=%d) %d %d\n", i+1, n, k, h, l)
	}
}

// seemed maybe closer.. but still wrong.
func solve2(n, k int) (h, l int) {
	h, l = n, n
	// height := level(k)
	kCurrent := 0
	depth := 0
	for kCurrent < k {
		h = h/2 + h%2
		l = l / 2
		kLevel := pow2(depth)
		if k-kCurrent < kLevel {
			fmt.Println("k-kCurrent<kLevel = ", kCurrent, kLevel, k-kCurrent)
			if kCurrent == int(pow2(depth+1)) {
				fmt.Println("doing first in level")
				h = h/2 + 1
				l = l / 2
			} else {
				fmt.Println("doing rest in row")
				h, l = h/2, l/2
			}
			kCurrent += kCurrent + kLevel
		} else {
			kCurrent -= kLevel
		}
		depth++
		fmt.Println(h, l, kCurrent, depth)
	}
	return
}

// i don't know why this isn't working. it seems to work for all the
// "explained" test cases (1-5) and some I made up.
//
// ... except it didnt: case 3 (n=6, k=2) was wrong. I'm getting (1,0), and the
// answer is (1,1). On paper I'm getting this of course, but my little "formula"
// isn't working for SOME cases when n is even (seems to work when n=1000...)
func solve(n, k int) (h, l int) {
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

	todepth := height(k)
	curdepth := 0
	visited := 0
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

func solve3(n, k int) (h, l int) {
	level := height(k)
	if level <= 0 {
		// base case
		return high(n), low(n)
	}

	// some redudancy for clarity
	lvlsq := pow2(level)
	start := lvlsq
	mid := lvlsq + lvlsq/2

	// do "branch" here
	var offset int
	var f func(int, int) int
	if k < mid {
		offset = k - start
		f = max
	} else {
		offset = k - mid
		f = min
	}

	// recurse to previous "level", get it's high and low,
	// and then choose the max or min (as f) as the n used to calculate
	// k's high and low
	kPrev := pow2(level-1) + offset
	nPrev := f(solve3(n, kPrev))

	return high(nPrev), low(nPrev)
}

// height of kth (1,2,...) stall taken
func height(k int) int {
	return int(math.Floor(math.Log2(float64(k))))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func pow2(x int) int {
	return 1 << uint(x)
}

func above(h int) int {
	return pow2(h+1) - 1
}

func high(n int) int {
	return n / 2
}

func low(n int) int {
	return (n - 1) / 2
}
