package game

import (
	"fmt"
	"math"
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
		fmt.Println(result)
	}
	fmt.Printf("%d çözüm kümesi bulundu\n", len(results))
}

func CreateTable(result []int) {
	for x := 0; x <= 8; x++ {

		for y := 0; y <= 8; y++ {

		}

	}
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
