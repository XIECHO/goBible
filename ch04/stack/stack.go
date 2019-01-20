package main

import "fmt"

func main() {
	stack := []int{1, 2}
	fmt.Println(stack)
	stack = append(stack, 3)
	fmt.Println(stack)
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
