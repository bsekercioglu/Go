package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift Sayılar Çalışıyor")
		time.Sleep(time.Nanosecond)
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek Sayılar Çalışıyor")
		time.Sleep(time.Nanosecond)
	}
	tekSayiCn <- toplam
}
