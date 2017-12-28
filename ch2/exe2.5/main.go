package main

import "fmt"

func PopCount(x uint64) int {
	var sum int = 0
	for x != 0 {
		x = (x & (x - 1))
		sum++
	}
	return sum
}
func main() {
	var x uint64 = 0x9876123450fedabc
	fmt.Printf("PopCount(0x%x) = %v\n", x, PopCount(x))
}
