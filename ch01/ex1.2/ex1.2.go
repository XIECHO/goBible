package main

import (
	"fmt"
	"os"
)

func main() {
	for i, item := range os.Args[1:] {
		fmt.Println(i, item)
	}
}
