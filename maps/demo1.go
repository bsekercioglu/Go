package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["desk"] = "sıra"
	sozluk["pencil"] = "kalem"
	sozluk["chair"] = "sandalye"
	fmt.Println(sozluk["desk"])
	fmt.Println("Sözlük: ", sozluk)

	fmt.Println("Eleman sayısı :", len(sozluk))
	delete(sozluk, "desk") //desk anahtarını ve değerini siliyoruz
	fmt.Println("Eleman sayısı :", len(sozluk))

}
