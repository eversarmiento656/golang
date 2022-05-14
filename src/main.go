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

type figuras interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
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

	//stringers
	myPc2 := pc{
		ram:   16,
		disk:  1000,
		marca: "msi",
	}
	fmt.Println(myPc2)

	// area con strunct
	c := cuadrado{base: 12}
	calcular(c)

	r := rectangulo{
		base:   12,
		altura: 3,
	}
	calcular(r)

	//lista de interfaces
	myInterface := []interface{}{
		"Hola", 12, true, 4.90}

	fmt.Println(myInterface)

}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.marca)
}

func (myPC pc) ping() {
	fmt.Println(myPC.marca, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras) {
	fmt.Println(f.area())
}
