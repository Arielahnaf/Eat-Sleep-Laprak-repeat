package main

import "fmt"

func main() {
	var f float64
	fmt.Scanln(&f)
	c := (f - 32) * 5 / 9

	fmt.Printf("Suhu dalam celcius: %2f/n", c)
}
