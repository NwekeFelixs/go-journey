package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	stop:= 0.0000001
	for {
		prevZ := z
		z -= (z*z -x) / (2*z)
		fmt.Println(z)
		if math.Abs(z - prevZ) < stop {
			break
		}
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
