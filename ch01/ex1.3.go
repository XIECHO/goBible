package main

import (
	"fmt"
	"strings"
	"time"
)

var args = []string{"hi", "there", "bubby", "boy", "5", "6", "7", "8", "9"}

func main() {
	start := time.Now()
	for i := 0; i < 30000; i++ {
		strings.Join(args, " ")
	}
	fmt.Println(time.Since(start).Seconds())
	start = time.Now()
	for i := 0; i < 30000; i++ {
		r, sep := "", ""
		for _, a := range args {
			r += sep + a
			sep = " "
		}
	}
	fmt.Println(time.Since(start).Seconds())
}
