package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12

}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12

}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 { //kişinin birden fazla kredisi olabileceği için interface yapımız array olarak gönderildi
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + float64(credits[i].Calculate())
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Samsundaki ev"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "Ankaradaki ev"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Polo Araba"}
	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam Kredi Ödemeniz :", total)
}
