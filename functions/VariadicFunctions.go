package functions

import "fmt"

func ToplaVariadic(sayilar ...int) int {

	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]

	}
	return toplam
}
func Variadic() {
	t := ToplaVariadic(5, 3, 65, 32, 54, 76, 3, 12, 54, 76, 76, 23)
	fmt.Printf("Toplam : %v\n", t)

}
