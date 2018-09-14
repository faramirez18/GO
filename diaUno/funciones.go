package main

import "fmt"

/*
func Sumar(x int, y int)  int {
	return x + y
}
*/
func main()  {
	//fmt.Println("La sumatoria es: ",Sumar(10,5))

	i := 42
	p := &i

	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}