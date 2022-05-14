package main

import (
	"curso_platzi/src/mypackage"
	"fmt"
)

const HOLA = "Saludo"

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 1000
	fmt.Println(myCar)

	mypackage.PrintMessage(HOLA)
}
