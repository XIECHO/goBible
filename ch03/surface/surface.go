package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
	fmt.Printf("test = %g\n", 5419816.66615614918189)
	fmt.Printf("test = %f\n", 5419816.66615614918189)
	fmt.Printf("test = %e\n", 5419816.66615614918189)

	fmt.Println(math.Inf(1) == math.Inf(2))

}
