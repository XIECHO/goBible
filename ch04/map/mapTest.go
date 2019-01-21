package main

import "fmt"

func main() {
	ages := make(map[string]int)
	sex := map[string]string{"bob": "female"}
	loc := map[string]string{
		"bob": "sh",
	}
	fmt.Println(ages, sex, loc)
	delete(sex, "bob")
	fmt.Println(sex)
	fmt.Println(&sex)

	t := 20
	fmt.Println(&t)
}
