package main

import (
	"fmt"
	"github.com/zhangbo1882/gopl-exercise/ch2/exe2.2/convert"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Input your digital number\n")
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v C = %v F\n", value, convert.CToF(convert.Celsius(value)))
		fmt.Printf("%v F = %v C\n", value, convert.FToC(convert.Fahrenhert(value)))
		fmt.Printf("%v feet  = %v meters\n", value, convert.FeetToMeters(convert.Feet(value)))
		fmt.Printf("%v meters = %v feet\n", value, convert.MetersToFeet(convert.Meters(value)))
		fmt.Printf("%v pounds = %v kilograms\n", value, convert.PoundsToKilo(convert.Pounds(value)))
		fmt.Printf("%v kilograms = %v pounds\n", value, convert.KiloToPounds(convert.Kilograms(value)))
	}
}
