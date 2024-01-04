package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtracao => ", a-b)
	fmt.Println("Multiplicacao => ", a*b)
	fmt.Println("Divisao => ", a/b)
	fmt.Println("Modulo => ", a%b)

	// Bit a bit (lógica)

	fmt.Println("And => ", a&b) // 11 & 10 = 10
	fmt.Println("Or => ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	// Math

	c := 3.0
	d := 2.0

	fmt.Println("Maior => ", math.Max(c, d)) // math tem que ser float64
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Potencia => ", math.Pow(c, d))
}
