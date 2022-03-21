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

		CreateTable(result)

	}
	fmt.Printf("%d çözüm kümesi bulundu\n", len(results))
}

func CreateTable(result []Point) {
	var tablo [8][8]string
	fmt.Println(result)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {

			tablo[x][y] = "."

		}

	}

	for item := 0; item < 8; item++ {
		item_x := result[item].x
		item_y := result[item].y

		for x := 0; x < 8; x++ {
			for y := 0; y < 8; y++ {
				if x == item_x && y == item_y {
					tablo[x][y] = "Q"
				}

			}

		}
	}

	for i := 0; i < 8; i++ {
		for w := 0; w < 8; w++ {

			fmt.Print(tablo[i][w])
		}
		fmt.Print("\n")
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
