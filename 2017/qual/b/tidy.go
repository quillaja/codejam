// https://code.google.com/codejam/contest/3264486/dashboard#s=p1

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var t int
	fmt.Scanln(&t)

	var N uint64
	for i := 0; i < t; i++ {
		fmt.Scanln(&N)
		result := solve(N)
		fmt.Printf("Case #%d: %d\n", i+1, result)
	}
}

func solve(n uint64) uint64 {
	for ; n > 0; n-- {
		b := []byte(strconv.FormatUint(n, 10))
		// fmt.Println(n, n/uint64(math.Pow10(len(b)-2)))
		if sort.SliceIsSorted(b, func(i, j int) bool {
			return b[i] < b[j]
		}) {
			return n
		}
	}
	return n
}
