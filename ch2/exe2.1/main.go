package main

import (
	"fmt"
	"github.com/zhangbo1882/gopl-exercise/ch2/exe2.1/tempconv"
)

func main() {
	var k tempconv.Kelvin = 100
	var c tempconv.Celsius = 40
	fmt.Printf("%v°K = %v°C\n", k, tempconv.KToC(k))
	fmt.Printf("%v°C = %v°K\n", c, tempconv.CToK(c))
}
