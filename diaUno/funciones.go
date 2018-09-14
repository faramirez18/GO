package main

import "fmt"

func Sumar(x int, y int)  int {
	return x + y
}

func main()  {
	fmt.Println("La sumatoria es: ",Sumar(10,5))
}