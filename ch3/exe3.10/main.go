package main

import "fmt"
import "bytes"

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

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	m := n % 3
	if m != 0 {
		//	fmt.Fprintf(&buf, "%s", s[:m])
		buf.WriteString(s[:m])
		buf.WriteByte(',')
	}
	for m < n {
		//	fmt.Fprintf(&buf, "%s", s[m:m+3])
		buf.WriteString(s[m : m+3])
		m += 3
		if m < n-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
func main() {
	s := "93834934389218123"
	fmt.Println(comma(s))
	fmt.Println(comma1(s))
	fmt.Println(comma2(s))
}
