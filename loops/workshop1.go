package loops

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Demo3() {

	fmt.Printf("Ortalama Nokta: %v\n", GetMedian(100, 253))
	rand.Seed(time.Now().UnixNano())
	min, max, count := 0, 1000, 0
	random := randInt(min, max)

	fmt.Printf("rastgele sayı:%v\n", random)
	fmt.Println("**************************************************************")
	fmt.Println("            Rastgele Sayı Üreterek Sayı Tahmini")

BASLA:
	number := randInt(min, max)

	if number < random {
		fmt.Printf("%v değerinden daha Büyük bir sayı tahmin ediniz (%v , %v arasında)\n", number, min, max)
		min = number
		count++
		goto BASLA
	} else if number > random {
		fmt.Printf("%v değerinden daha Küçük bir sayı tahmin ediniz (%v , %v arasında)\n", number, min, max)
		max = number
		count++
		goto BASLA
	} else {
		fmt.Printf("%v tahmininizi %v denemede doğru bildiniz\n", number, count)
	}

	fmt.Println("**************************************************************")
	fmt.Println("            İkiye Bölme Yöntemi ile Sayı Tahmini")

	min, max, count, number = 0, 1000, 0, 0
BASLA2:
	number = int(GetMedian(float64(min), float64(max)))

	if number < random {
		fmt.Printf("%v değerinden daha Büyük bir sayı tahmin ediniz (%v , %v arasında)\n", number, min, max)
		min = number
		count++
		goto BASLA2
	} else if number > random {
		fmt.Printf("%v değerinden daha Küçük bir sayı tahmin ediniz (%v , %v arasında)\n", number, min, max)
		max = number
		count++
		goto BASLA2
	} else {
		fmt.Printf("İkiye bölme yöntemi ile %v tahmininizi %v denemede doğru bildiniz\n", number, count)
	}

}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)

}

func GetMedian(min float64, max float64) float64 {
	total := min + max
	return math.Round(total / 2)

}
