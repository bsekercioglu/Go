package maps

import "fmt"

func Demo2() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range sozluk {
		fmt.Println(k + "\t" + v)

	}

}
