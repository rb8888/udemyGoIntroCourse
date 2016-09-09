// Here's my answer for the exercise at:
// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Sqrt(x float64) float64 {
	var z float64 = rand.Float64()
	for i:=0; i < 10; i++ {
		zn := z - ((math.Pow(z,2) - x)/(2*z))
		if math.Abs(z-zn) < 0.01 {
			return zn
		}
		z = zn
	}
	return z
}

func main() {
	fmt.Println("Newton's method =", Sqrt(2))
	fmt.Println("math.Sqrt =", math.Sqrt(2))
}