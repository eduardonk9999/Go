package main

import "fmt"

func main() {
	nome := "Eduardo"
	idade := 34

	println(nome, idade)

	//constantes...
	const PI = 3.14159

	// Conversão de tipos
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	println(i, f, u)

	var preco float64 = 19.90
	println(preco)

	const MAX_USUARIOS = 100
	var usuarios = 0

	print(usuarios)

	var valorInicial int = 4

	var valorConvertido float64 = float64(valorInicial) / 2

	fmt.Printf("%.2f\n", valorConvertido)
}
