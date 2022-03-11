package main

import (
	"golesson/conditionals"
	"golesson/loops"
	"golesson/variables"
)

/*
Açıklama: Yeni modül oluşturmak için terminal ekranında "go mod init golesson" yazılır.
Bunun anlamı golesson modülü yaratıyoruz.
ve modülü kullanmak istediğimiz kısımda ise import olarak bu modul ve ona bağlı altklasör bilgisi verilir
"golesson/variables" variables klasöründeki modülü kullanacağız...

mdoülü kullanırken de variables altındaki Demo1 metodunu kullanacağız şeklindedir.
variables.Demo1()



*/
func main() {
	conditionals.Demo1()
	variables.Demo1()
	conditionals.Demo2()
	loops.Demo1()
	loops.Demo3()

}
