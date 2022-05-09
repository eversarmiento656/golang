package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base: ", base)
	fmt.Println("altura: ", altura)
	fmt.Println("area: ", area)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado: ", areaCuadrado)

	// Operadores aritmeticos
	x := 10
	y := 20
	suma := x + y
	resta := x - y
	multiplicacion := x * y
	division := x / y

	fmt.Println("suma: ", suma)
	fmt.Println("resta: ", resta)
	fmt.Println("multiplicacion: ", multiplicacion)
	fmt.Println("division: ", division)
	fmt.Println("Modulo: ", x%y)

	// Uso de fmt

	var hello string = "hello"
	var world string = "world"
	fmt.Println(hello, world)

	// Printf
	meses := "Meses"
	numMeses := 12
	fmt.Printf("el año tiene %d %s\n", numMeses, meses)
	fmt.Printf("el año tiene %v %v\n", numMeses, meses)

	// SPrintf
	message := fmt.Sprintf("el año tiene %d %s", numMeses, meses)
	fmt.Println(message)

	//Imprimir tipo de dato
	fmt.Printf("meses: %T\n", meses)
	fmt.Printf("numMeses: %T\n", numMeses)

}
