package pointers

import "fmt"

func Demo1(sayi int) {

	sayi = sayi + 1
	fmt.Println("Demo : ", sayi)
}

func Demo2(sayi *int) {

	*sayi = *sayi + 1
	fmt.Println("Demo2 : ", sayi)
}
