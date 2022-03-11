package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 10000
	var cekilmekIstenen float64 = 900
	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersiz : " + fmt.Sprintf("<%f>\n", hesap)) //Sprintf ile formatlı yazdırma yapılabilir.
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız Hazırlanıyor : ")
		fmt.Println("Dikkat, hesapta para kalmadı")
	} else {
		fmt.Println("Paranız hazırlanıyor")

	}

}
