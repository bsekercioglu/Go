package loops

//kullanıcıdan sayık girmesini isteyin
//23: Asalsayıdır
import "fmt"

func Demo4() {

	sayi := 0
	fmt.Print("Lütfen kontrol edilmesini istediğiniz sayıyı girin:")
	fmt.Scanf("%d", &sayi)

	if sayi%2 == 0 || sayi%3 == 0 || sayi%5 == 0 || sayi%7 == 0 {
		fmt.Printf("%v sayısı Asal Sayı DEĞİLDİR\n", sayi)
	} else {
		fmt.Printf("%v sayısı Asal SAYI\n", sayi)
	}

}
