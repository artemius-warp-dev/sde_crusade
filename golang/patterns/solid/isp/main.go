package main

import (
	"fmt"
	"math"
)

// Bad code

// type square struct {
// 	length float64
// }

// func (s square) area() float64 {
// 	return s.length * s.length
// }

// // square is a flat surface -> does not have volume
// func (s square) volume() float64 {
// 	return 0
// }

// type cube struct {
// 	length float64
// }

// func (c cube) area() float64 {
// 	return math.Pow(c.length, 2)
// }

// func (c cube) volume() float64 {
// 	return math.Pow(c.length, 3)
// }

// type shape interface {
// 	area() float64
// 	volume() float64
// }

// func areaSum(shapes ...shape) float64 {
// 	var sum float64
// 	for _, s := range shapes {
// 		sum += s.area()
// 	}
// 	return sum
// }

// func areaVolumeSum(shapes ...shape) float64 {
// 	var sum float64
// 	for _, s := range shapes {
// 		sum += s.area() + s.volume()
// 	}
// 	return sum
// }

// func main() {
// 	s := square{length: 3}
// 	c := cube{length: 4}
// 	fmt.Println(areaSum(s, c))
// 	fmt.Println(areaVolumeSum(s, c))
// }

//Good code

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s1 := square{length: 5}
	s2 := square{length: 6}
	c1 := cube{square: square{length: 3}}
	c2 := cube{square: square{length: 4}}
	fmt.Println(areaSum(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum(c1, c2))
}
