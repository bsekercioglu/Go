package loops

import (
	"fmt"
	"math/rand"
	"time"
)

func Demo3() {

	rand.Seed(time.Now().UnixNano())
	random := randInt(0, 100)
	min, max, count := 0, 100, 0
	fmt.Printf("rastgele sayı:%v\n", random)
BASLA:
	number := randInt(min, max)

	if number < random {
		fmt.Printf("%v değerinden daha Büyük bir sayı tahmin ediniz\n", number)
		min = number
		count++
		goto BASLA
	} else if number > random {
		fmt.Printf("%v değerinden daha Küçük bir sayı tahmin ediniz\n", number)
		max = number
		count++
		goto BASLA
	} else {
		fmt.Printf("%v tahmininizi %v denemede doğru bildiniz\n", number, count)
	}

}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)

}
