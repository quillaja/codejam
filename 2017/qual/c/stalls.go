// https://code.google.com/codejam/contest/dashboard?c=3264486#s=p2
// first impression is that this is a binary tree type problem
// but with N up to 10^18, this can't store "state" (it's an exabit
// of memory at a minimum)
//
// last impressions: this was pretty hard for me, and thinking about it as a tree
// didn't help--maybe hurt--although it did lead me to useful formulas regarding
// binary trees. Inclusion of buggy numbers (see comments on math.Log2) is sneaky.

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var t int
	fmt.Scanln(&t)
	// t = 15
	var n, k int
	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		h, l := solve(n, k)
		fmt.Printf("Case #%d: %d %d\n", i+1, h, l)
		// fmt.Printf("Case #%d:(n=%d k=%d) %d %d\n", i+1, n, k, h, l)
	}
}

// Recursive solution using some properties of binary trees (height, nodes at
// a particular height, etc). It starts at the "bottom" (level of k), and proceeds
// up in levels (height) until reaching the root. At each level, it calculates
// which of the previous 2^(lvl-1) nodes is it's "parent", and if it should use
// the high or low value of the parent (max or min of high/low). Nodes in the
// first half of a particular level will always take the high value of their
// parent, and nodes in the second half of their level will take the low value
// of their parent.
// Since k is in [1, 10^18], the max hight and therefore max recursion depth
// should be about 60 (log2(1e18) = 59.79...).
func solve(n, k int) (h, l int) {
	level := height(k)
	if level <= 0 {
		// base case
		h, l = high(n), low(n)
		// fmt.Printf("lv=%d (%d, %d) k=%d | ", level, h, l, k)
		return
	}

	// some redudancy for clarity
	twoPowLvl := pow2(level)
	start := twoPowLvl
	mid := twoPowLvl + twoPowLvl/2

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

	// for debugging math.Log2 error. see height() for explanation.
	// if offset < 0 {
	// 	panic(fmt.Errorf("offset <0: offset=%d\nlevel=%d k=%d\nstart=%d mid=%d\ntwoPowLvl=%d",
	// 		offset, level, k, start, mid, twoPowLvl))
	// }

	// recurse to previous "level", get it's high and low,
	// and then choose the max or min (as f) as the n used to calculate
	// k's high and low
	kPrev := pow2(level-1) + offset
	nPrev := f(solve(n, kPrev))

	h, l = high(nPrev), low(nPrev)
	// fmt.Printf("lv=%d (%d, %d) k=%d kP=%d | ", level, h, l, k, kPrev)
	return
}

// height of kth (1,2,...) stall taken
func height(k int) int {
	// CANNOT use math.Log2 function. Using floats will give the wrong answer
	// for 288230376151711743 and probably others because the precision
	// of float64 causes 57.99999999999999999499464608778919110696411496790292147213...
	// to round up to 58.
	// return int(math.Floor(math.Log2(float64(k))))
	return log2Fast(k)
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

func high(n int) int {
	return n / 2
}

func low(n int) int {
	return (n - 1) / 2
}

func log2Slow(x int) int {
	i := 0
	for ; x > 1; i++ {
		x = x >> 1
	}
	return i
}

func log2Fast(x int) int {
	// for integers, log2() is essentially finding the number of bits
	// required to store the number
	return bits.Len64(uint64(x)) - 1
}
