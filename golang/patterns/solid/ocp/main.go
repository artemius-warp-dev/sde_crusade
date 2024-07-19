package main

//Bad code

import (
	"fmt"
	"math"
)

// type circle struct {
// 	radius float64
// }

// type square struct {
// 	length float64
// }

// type calculator struct {
// }

// func (a calculator) areaSum(shapes ...interface{}) float64 {
// 	var sum float64
// 	for _, shape := range shapes {
// 		switch shape.(type) {
// 		case circle:
// 			r := shape.(circle).radius
// 			sum += math.Pi * r * r
// 		case square:
// 			l := shape.(square).length
// 			sum += l * l
// 		}
// 	}
// 	return sum
// }

// func main() {
// 	c := circle{radius: 5}
// 	s := square{length: 7}
// 	calc := calculator{}
// 	fmt.Println("area sum:", calc.areaSum(c, s))
// }

//Good code

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

type calculator struct {
}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	t := triangle{height: 3, base: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s, t))
}
