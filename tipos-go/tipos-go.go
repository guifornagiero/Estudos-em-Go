package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a := 5
	fmt.Println("Inteiro: ", reflect.TypeOf(a))

	c := math.MaxInt64
	fmt.Println("Int64: ", reflect.TypeOf(c), c)

	var b byte = 255
	fmt.Println("Byte (uint8): ", reflect.TypeOf(b))

	var d rune = 'a' // rune = char
	fmt.Println("Rune (ASCII): ", reflect.TypeOf(d), d)

	e := 41.23
	fmt.Println("Float: ", reflect.TypeOf(e))

	s1 := "olá meu nome é guilherme"
	fmt.Println("String: ", reflect.TypeOf(s1), len(s1))

	s2 := `olá
	meu
	nome
	é
	guilherme`

	fmt.Println("String (quebra de linha): ", reflect.TypeOf(s2), len(s2))

}