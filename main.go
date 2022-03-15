package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	functions.Variadic()
	sayilar := []int{5, 3, 65, 32, 7887, 54, 76, 3, 12, 54, 76, 76, 23}
	fmt.Printf("Toplam : %v\n", functions.ToplaVariadic(sayilar...))
}
