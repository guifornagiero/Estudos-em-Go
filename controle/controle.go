package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.Intn(10)
	return s
}

// Switch de tipos
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "funçao"
	case rune:
		return "rune"
	default:
		return "nao sei"
	}
}

func main() {
	// If com inicialização de variável na declaração
	if i := randomNumber(); i > 5 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	// For com o mesmo funcionamento do while
	j := 0
	for j <= 10 {
		fmt.Printf("%d ", j)
		j++
	}

	fmt.Println()

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// for {
	// 	// Laço infinito...
	// }

	fmt.Println()

	// Switch padrão - comparação de casos
	nota := 6
	switch nota {
	case 10:
		fallthrough // Desce para o case de baixo
	case 9:
		fmt.Println("A")
	case 8: // Assim NÃO desce para o de baixo... Não é igual C#
	case 7:
		fmt.Println("B")
	case 6, 5: // Desce para o de baixo
		fmt.Println("C")
	default:
		fmt.Println("Nota invalida")
	}

	// Switch true - busca o primeiro resultado que seja true
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Bom dia")
	case time.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	fmt.Println(tipo(3))
	fmt.Println(tipo(3.3))
	fmt.Println(tipo("string"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo('a'))

}
