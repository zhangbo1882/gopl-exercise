package main

import "fmt"
import "os"

func IsAnagrams(s string, t string) bool {
	if s == t {
		return false
	}
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}
	mapS := make(map[byte]int)
	mapT := make(map[byte]int)
	for i := 0; i < lenS; i++ {
		_, ok := mapS[s[i]]
		if !ok {
			mapS[s[i]] = 1
		} else {
			mapS[s[i]]++
		}
		_, ok = mapT[t[i]]
		if !ok {
			mapT[t[i]] = 1
		} else {
			mapT[t[i]]++
		}
	}
	for name, counter := range mapS {
		if mapT[name] != counter {
			return false
		}
	}
	return true

}
func main() {
	var b bool
	if len(os.Args) != 3 {
		fmt.Println("Input two strings")
		os.Exit(1)
	}
	var s1 string = os.Args[1]
	var s2 string = os.Args[2]
	b = IsAnagrams(s1, s2)
	if b == true {
		fmt.Printf("%s and %s are  anagrams\n", s1, s2)
	} else {
		fmt.Printf("%s and %s are  not anagrams\n", s1, s2)
	}

}
