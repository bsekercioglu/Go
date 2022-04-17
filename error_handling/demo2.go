package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}
	return "Bildiniz", nil
}

func Demo2() {
	msj, _ := TahminEt(50)
	fmt.Println(msj)
	fmt.Println(TahminEt(0))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-50))

}
