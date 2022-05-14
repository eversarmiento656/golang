package main

import (
	"curso_platzi/src/mypackage"
	"fmt"
)

const HOLA = "Saludo"

type pc struct {
	ram   int
	disk  int
	marca string
}

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 1000
	fmt.Println(myCar)

	mypackage.PrintMessage(HOLA)

	//punteros
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{
		ram:   6,
		disk:  200,
		marca: "msi",
	}

	fmt.Println(myPc)
	myPc.ping()
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}

func (myPC pc) ping() {
	fmt.Println(myPC.marca, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}
