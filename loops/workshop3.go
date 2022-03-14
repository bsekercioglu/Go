package loops

import "fmt"

func ArkadasSayilar() {
	maxSayi := 10000
	for j := 1; j < maxSayi; j++ {

		durum, sayi := ArkadasSayiKontrol(j)
		if durum == true {
			fmt.Printf("%v sayısı %v ile Arkadaş Sayıdır \n", j, sayi)
		}
	}

}

func ArkadasSayiKontrol(sayi int) (bool, int) {

	toplam, toplam2 := 0, 0

	for i := 1; i < sayi; i++ {

		if sayi%i == 0 {
			toplam = toplam + i

		}
	}
	for j := 1; j < toplam; j++ {
		if toplam%j == 0 {
			toplam2 = toplam2 + j

		}
	}
	if sayi == toplam2 && sayi != toplam {
		return true, toplam
	} else {
		return false, toplam
	}

}
