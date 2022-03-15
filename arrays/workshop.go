package arrays

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateArray(size int) ([]int, time.Duration) {
	start := time.Now()
	var sayilar []int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		rnd := randInt(0, size)
		sayilar = append(sayilar, rnd)
	}
	elapsed := time.Since(start)

	return sayilar, elapsed
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	eleman := 500000
	dizi, sure := CreateArray(eleman)
	//fmt.Println(dizi)
	fmt.Printf("%v Elemanlı dizi oluşturma süresi : %v\n", eleman, sure)
	fmt.Println("*******************************************************")
	dizi, sure = SelectionSort(dizi)
	fmt.Printf("Selection Sort (Süre:%v): \n", sure)
	fmt.Println("*******************************************************")
	dizi, sure = InsertionSort(dizi)

	fmt.Printf("Insertion Sort (Süre:%v): \n", sure)
	fmt.Println("*******************************************************")
}

func SelectionSort(liste []int) ([]int, time.Duration) {
	start := time.Now()
	for i := 0; i < len(liste); i++ {

		for j := i + 1; j < len(liste); j++ {
			if liste[j] < liste[i] {

				liste[i], liste[j] = liste[j], liste[i] //dizi elemanlarına swap uyguluyoruz

			}

		}

	}
	elapsed := time.Since(start)

	return liste, elapsed
}

func InsertionSort(liste []int) ([]int, time.Duration) {
	start := time.Now()
	for i := 0; i < len(liste); i++ {

		value := liste[i]
		j := i - 1

		for j >= 0 && liste[j] > value {
			liste[j+1] = liste[j]
			j = j - 1
		}
		liste[j+1] = value

	}
	elapsed := time.Since(start)

	return liste, elapsed
}
