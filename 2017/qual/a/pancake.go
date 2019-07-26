// https://code.google.com/codejam/contest/3264486/dashboard#s=p0

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var t int
	fmt.Scanln(&t)

	pcakes := []byte{}
	flipper := 0
	for i := 0; i < t; i++ {
		fmt.Scanln(&pcakes, &flipper)
		result := solve(pcakes, flipper)
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
}

// this just goes through the string, flips whenever a minus appears, and
// keeps going until the end of the string. if every char in the string is not
// a plus after this, it is "impossible".
//
// note: i seriously cannot believe the solution to this problems was this
// damn simple. i didn't expect this to work--but it magically did.
func solve(pcakes []byte, flipper int) string {
	flips := 0
	for i := 0; i <= len(pcakes)-flipper; i++ {
		if pcakes[i] == minus {
			flip(pcakes, i, flipper)
			flips++
		}
	}
	if countPlus(pcakes) != len(pcakes) {
		return "IMPOSSIBLE"
	}
	return strconv.Itoa(flips)
}

const (
	plus  = '+'
	minus = '-'
)

func flip(pcakes []byte, start, flipper int) {
	if start < 0 || start+flipper > len(pcakes) {
		return // invalid flipping params
	}

	// reverse all values from start to the 'end' of the flipper
	for i := start; i < start+flipper; i++ {
		switch pcakes[i] {
		case plus:
			pcakes[i] = minus
		case minus:
			pcakes[i] = plus
		}
	}
}

func countPlus(pcakes []byte) int {
	return bytes.Count(pcakes, []byte{plus})
}
