package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo2.txt")

	//Dosya bulunursa err=nil , f=dosyaadi olur
	if err != nil {
		//type assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya Bulunamad─▒ :", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamad─▒ :")
			return
		}

	} else {
		fmt.Println("Dosya bulundu: ", f.Name())
	}
}
