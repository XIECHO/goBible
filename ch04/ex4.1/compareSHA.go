package main

var pc = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func compare(x, y []byte) int {
	var sum int
	temp = x & y
	sum 
}
func main() {
  
}
