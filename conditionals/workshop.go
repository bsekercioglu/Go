package conditionals

import "fmt"

func Demo2() {
	//3 adet int değişken tanımla
	//Ekrana en büyük olanı yazdır.

	var sayi1, sayi2, sayi3 int = 50, 50, 30

	var enbuyuk string

	if sayi1 > sayi2 {

		if sayi1 > sayi3 {
			enbuyuk = "sayi1"

		} else {
			enbuyuk = "sayi3"
		}
	} else if sayi2 > sayi1 {
		if sayi2 > sayi3 {
			enbuyuk = "sayi2"
		} else {
			enbuyuk = "sayi3"
		}

	}
	fmt.Println(enbuyuk)
}
