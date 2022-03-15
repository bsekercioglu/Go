package slices

import "fmt"

func AppendSlice() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "Burak"
	isimler[1] = "Ata"
	isimler[2] = "Ada"
	/*Yeni bir eleman eklerken APPEND kullanılır, yoksa hata oluşur*/
	isimler = append(isimler, "Tuba")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
func CopySlice() {
	sehirler := []string{"Samsun", "İstanbul", "Ankara", "İzmir"}
	fmt.Println(sehirler)
	SehirlerKopya := make([]string, len(sehirler))
	fmt.Println(SehirlerKopya)
	copy(SehirlerKopya, sehirler)
	sehirler = append(sehirler, "Rize")
	fmt.Println(sehirler)
	fmt.Println(SehirlerKopya)

	fmt.Println(sehirler[1:])  //1. indisden sonuna kadar yazdır
	fmt.Println(sehirler[1:3]) //1-3. indis arasını yazdır
	fmt.Println(sehirler[:2])  //2. indise kadar yazdır

}
