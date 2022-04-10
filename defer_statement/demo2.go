package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift Sayı"
	}
	if sayi%2 != 0 {
		return "Tek Sayı"
	}

	return "Belirsiz"
}

func Test() {
	fmt.Println(Demo2(5))
}

func DeferFunc() {
	fmt.Println("Bitti")
}
