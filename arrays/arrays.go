package arrays

import "fmt"

func Demo6() {
	var sayilar [5]int
	sayilar[2] = 50
	fmt.Println(sayilar)
	for i := 0; i < 5; i++ {
		fmt.Printf("%v. sayi %v\n", i, sayilar[i])
	}

}
