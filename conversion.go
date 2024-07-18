package main

import (
	"fmt"
	"math"
)

func c2f(c int) float64 {
	return float64(c)*9/5 + 32
}

func isMirror(a int, b int) bool {
	return reverseString(massage(a)) == massage(b)
}

func massage(n int) string {
	switch {
	case n < 10:
		return fmt.Sprintf("0%d", n)
	case n >= 100:
		return massage(n - 100)
	default:
		return fmt.Sprintf("%d", n)
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func printConversion(c int, f int) {
	fmt.Println(fmt.Sprintf("%d°C ~= %d°F", c, f))
}

func main() {
	for c := 4; c < 100; c += 12 {
		var f = c2f(c)
		if isMirror(c, int(math.Ceil(f))) {
			printConversion(c, int(math.Ceil(f)))
		} else if isMirror(c, int(math.Floor(f))) {
			printConversion(c, int(math.Floor(f)))
		} else {
			break
		}
	}
}
