package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	for n > 3 {
		s = s[:n-3] + "," + s[n-3:]
		n -= 3
	}
	return s
}
func main() {
	s := "93834934389218123"
	fmt.Println(comma(s))
	fmt.Println(comma1(s))
}
