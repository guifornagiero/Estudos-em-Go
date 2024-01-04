package main

import (
	"fmt"
	m "math" // usar 'm' ao invés de 'math'
)

func main() {
	fmt.Println("Hello world!")

	const PI float64 = 3.14         	// definição direta de constante
	var raio float64 = 3.0          	// definição direta de variável (tem que ser usada)

	area := PI * m.Pow(raio, 2)     	// definição direta de variável sem informar o tipo (faz automaticamente)
	fmt.Println("Area: ", area)

	const (                         	// definição de constantes em bloco
		a = 1
		b = 2
	)

	var (                           	// definição de variáveis em bloco
		c = 3
		d = 4
	)

	var e, f bool = true, false     	// definição com tipo em mais de uma
	g, h, i := 2, false, "guilherme" 	// definição sem tipo em mais de uma

	fmt.Println(c, d, e, f, g, h, i)

	x := 3.1515
	xs := fmt.Sprint(x)					// converter float64 em string
	fmt.Println("O valor de x em string: " + xs)

	//  %d = int
	//  %f = float
	//  %s = string
	//  %t = boolean
	//  %v = todos
}