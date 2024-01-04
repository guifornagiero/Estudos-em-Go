package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i
	*p++
	i++

	fmt.Println(p, *p, i, &i)

	// p = endereço de memória que o 'p' aponta
	// *p = valor dentro desse endereço de memória
	// i = valor de i
	// &i = endereço de memória da variável i

}
