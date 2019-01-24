package main

import "fmt"

func main() {
	a := map[string]string{}
	fmt.Println(a)
	ni(a)
	fmt.Println(a)
	wo(&a)
	fmt.Println(a)

	b := []int{1, 2, 3}
	fmt.Println(b)
	ta(b)
	fmt.Println(b)

}

func ni(fir map[string]string) {
	fir["ni"] = "lll"
}

func wo(sec *map[string]string) {
	(*sec)["wo"] = "ggg"
}

func ta(thi []int) {
	thi[0] = 9
}
