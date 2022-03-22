package game

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x int
	y int
}

var results = make([][]Point, 0)

func Game() {
	Solve(8)

}

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
	fmt.Print("Sonuçlar:\n")
	for _, result := range results {

		CreateTable(result)

	}
	fmt.Printf("%d çözüm kümesi bulundu\n", len(results))
}

type Sorter []Point

func (a Sorter) Len() int           { return len(a) }
func (a Sorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sorter) Less(i, j int) bool { return a[i].x < a[j].x }

func CreateTable(result []Point) {
	var tablo [8][8]string
	sort.Sort(Sorter(result))

	fmt.Println(result)

	for x := 0; x < 8; x++ {
		sayac := x
		for y := 0; y < 8; y++ {
			sayac++
			if sayac%2 == 0 {
				tablo[x][y] = string(rune('\u25a0'))
			} else {
				tablo[x][y] = string(rune('\u25a1'))
			}

		}

	}

	for item := 0; item < 8; item++ {
		item_x := result[item].x
		item_y := result[item].y

		for x := 0; x < 8; x++ {
			for y := 0; y < 8; y++ {
				if x == item_x && y == item_y {
					tablo[x][y] = string(rune('\u265B'))
				}

			}

		}
	}
	fmt.Print("    ")
	for j := 0; j < 8; j++ {
		fmt.Printf(" %c ", rune(j+65))
	}
	fmt.Print("\n")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %v ", i+1)
		for w := 0; w < 8; w++ {

			fmt.Print("  ", tablo[i][w])
		}
		fmt.Print("\n")
	}
	fmt.Print("    ")
	for j := 0; j < 8; j++ {
		fmt.Printf(" %c ", rune(j+65))
	}
	fmt.Print("\n\n")
}

func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results = append(results, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)

				}

			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {

	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))

	return answer
}
