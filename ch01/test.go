package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("ni")
	temp := make(map[int]int)
	temp[1] = 10
	var a = 8
	var ok bool
	a, ok = temp[1]
	aa := []int{}
	aa = append(aa, 1)

	fmt.Println(aa)
	fmt.Println(a, ok)
	os.Exit(10)
	//	log.Fatal("kk")
}
