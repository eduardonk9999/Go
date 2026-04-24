package main

import "fmt"

func main() {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else if nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Enquanto
	for x < 100 {
		x *= 2
	}

	

}