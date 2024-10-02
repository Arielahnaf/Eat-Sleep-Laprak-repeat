package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64
	const pi = math.Pi

	fmt.Print("Jika jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	luas := pi * jariJari * jariJari

	fmt.Printf("Maka luas lingkaran dengan jari-jari %.2f adalah %.1f\n", jariJari, luas)
}
