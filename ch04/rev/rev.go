package main

import (
	"fmt"
	"sort"
)

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	i := []int{3, 76, 2}
	fmt.Println(reverse(i))
	sort.Ints(i)
	//sort.Reverse(i)
	fmt.Println(i)
}
