package main

import "fmt"

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c.Write([]byte("hellof"))
	fmt.Println(c)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
