package main

import (
	"fmt"
	"math"
)

/*Разработать программу нахождения расстояния между двумя точками,
  которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором. */

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func main() {
	p1 := NewPoint(2, 5)
	p2 := NewPoint(4, 8)

	distance := Distance(p1, p2)

	fmt.Printf("Расстояние между точками %v и %v = %.2f", p1, p2, distance)
}
