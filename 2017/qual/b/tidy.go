// https://code.google.com/codejam/contest/3264486/dashboard#s=p1

package main

import (
	"fmt"
	"math"
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

// break the number down, proceed to examine each digit. if a digit is less
// than the next one (ie it's not 'tidy'), "max" it and all previous digits
// by setting them to 9 and subtract 1 from the next digit.
func solve(n uint64) uint64 {
	b := deconstruct(n)
	for i := 0; i < len(b)-1; i++ {
		if b[i] < b[i+1] {
			for j := i; j >= 0; j-- {
				b[j] = 9
			}
			b[i+1]--
		}
	}
	return rebuild(b)
}

// break n into multiples of powers of 10, index 0 being the least significant
// digit.
func deconstruct(n uint64) []uint8 {
	p := []uint8{}
	for n != 0 {
		r := uint8(n % 10)
		n = n / 10
		p = append(p, r)
	}
	return p
}

// get back the number from the slice of digits.
func rebuild(p []uint8) uint64 {
	var num uint64
	for i, n := range p {
		num += uint64(n) * uint64(math.Pow10(i))
	}
	return num
}

// first brute force attempt. worked on small set but took WAY too long
// on larger numbers.
func solveSmall(n uint64) uint64 {
	for ; n > 0; n-- {
		b := []byte(strconv.FormatUint(n, 10))
		if sort.SliceIsSorted(b, func(i, j int) bool {
			return b[i] < b[j]
		}) {
			return n
		}
	}
	return n
}
