package mypackage

import "fmt"

//CarPublic si se coloca en mayuscula indicamos que la struct es publica
type CarPublic struct {
	Brand string
	Year  int
}

//CarPrivate si se coloca en minuscula indicamos que la struct es privado
type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(message string) {
	fmt.Println(message)
}
