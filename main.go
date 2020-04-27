package main

import (
	"fmt"
	"mehdisbys/justforfunc/merge"
)

func main() {
	a := merge.AsChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := merge.AsChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := merge.AsChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	for v := range merge.MergeRecursion(a, b, c) {
		fmt.Println(v)
	}
}

