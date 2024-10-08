package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(10000))
}
