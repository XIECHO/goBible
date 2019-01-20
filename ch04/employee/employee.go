package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Postion   string
	Salary    int
	ManagerID int
}

func main() {
	a := map[int]*Employee{}
	a[0] = &Employee{}
	a[0].Salary = 500

	b := map[int]Employee{}
	b[0] = Employee{}
	//b[0].Salary = 500

	var tmp = &[]Employee{Employee{}, Employee{Salary: 9000}}
	var temp = map[int]*Employee{}
	t := []int{}
	tt := []*int{}
	for i, item := range *tmp {
		fmt.Println(i, item)
		temp[i] = &item
		fmt.Println(temp)
		t = append(t, i)
		tt = append(tt, &i)
	}

	for k, v := range temp {
		fmt.Println(k, v)
	}
	fmt.Println(temp)
	fmt.Println(t)
	fmt.Println(tt)

	var ll Employee
	fmt.Println(ll)

	var lll map[string]int
	fmt.Println(lll)

}
