package interfaces

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r circle) area() float64 {
	return math.Pi * math.Pow(2, r.radius)

}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func Yazdir(s shape) {
	fmt.Printf("%v şeklinin Alanı %v cm \n", getType(s), s.area())

}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	Yazdir(r)
	r2 := circle{radius: 10}
	Yazdir(r2)
}

func getType(myvar interface{}) string {
	return reflect.TypeOf(myvar).Name() //myvar ile gelen interface bilgisinin ismini alıyoruz
}
