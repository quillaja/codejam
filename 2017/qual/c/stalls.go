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
	t = 15
	var n, k uint64
	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		h, l := solve(n, k)
		fmt.Printf("Case #%d: %d %d\n", i+1, h, l)
	}
}

// i don't know why this isn't working. it seems to work for all the
// "explained" test cases (1-5) and some I made up.
//
// ... except it didnt: case 3 (n=6, k=2) was wrong. I'm getting (1,0), and the
// answer is (1,1). On paper I'm getting this of course, but my little "formula"
// isn't working for SOME cases when n is even (seems to work when n=1000...)
func solve(n, k uint64) (h, l uint64) {
	// level := int(math.Ceil(math.Log2(float64(k+1)) - 1))
	level := int(math.Floor(math.Log2(float64(k))))
	// fmt.Println(level)
	h = n
	l = n
	for i := 0; i <= level; i++ {
		// for i := 0; k >= uint64(math.Pow(2, float64(i))); i++ {
		h = h / 2
		if l > 0 {
			l = (l - 1) / 2
		}
		// if h > 0 {
		// 	h = (h-1)/2 + 1
		// }
		// if l > 0 {
		// 	l = (l - 1) / 2
		// }
		fmt.Println(h, l)
		// if l > h {
		// l, h = h, l
		// }
	}

	return
}
