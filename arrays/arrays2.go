package arrays

import "fmt"

func Sort() {

	sayilar := [5]int{5, 10, 15, 22, 4}
	for i := 0; i < len(sayilar); i++ {
		for j := i + 1; j < len(sayilar)-1; j++ {
			if sayilar[i] > sayilar[j] {
				tmp := sayilar[i]
				sayilar[i] = sayilar[j]
				sayilar[j] = tmp
			} else if sayilar[i] < sayilar[j] {
				tmp := sayilar[j]
				sayilar[j] = sayilar[i]
				sayilar[i] = tmp
			}

		}

	}
	fmt.Println(sayilar)
}
