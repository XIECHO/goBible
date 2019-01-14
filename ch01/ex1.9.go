package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		fmt.Println(resp.Status)
		fmt.Print(strconv.Atoi("90"))
	}
}
