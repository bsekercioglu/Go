package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün Kaydedildi : ", p.productName)
	defer Log()
}
func Log() {
	fmt.Println("Loglandı")
}
func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "PC", unitPrice: 9000}

	fmt.Println("Bitti")
	fmt.Println(p.productName)
}
