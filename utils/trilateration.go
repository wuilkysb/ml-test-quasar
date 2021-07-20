package utils

import "math"

type Point struct {
	X float64
	Y float64
	R float64
}

type Solution []Point

func (s Solution) First() Point {
	return s[0]
}

func square(v float64) float64 {
	return v * v
}

func normalize(p Point) float64 {
	return math.Sqrt(square(p.X) + square(p.Y))
}

func dot(p1, p2 Point) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}

func subtract(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
	}
}

func add(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func divide(p Point, v float64) Point {
	return Point{
		X: p.X / v,
		Y: p.Y / v,
	}
}

func multiply(p Point, v float64) Point {
	return Point{
		X: p.X * v,
		Y: p.Y * v,
	}
}

func Solve(p1, p2, p3 Point) (Point, error) {
	ex := divide(subtract(p2, p1), normalize(subtract(p2, p1)))
	i := dot(ex, subtract(p3, p1))
	a := subtract(subtract(p3, p1), multiply(ex, i))
	ey := divide(a, normalize(a))
	d := normalize(subtract(p2, p1))
	j := dot(ey, subtract(p3, p1))

	x := (square(p1.R) - square(p2.R) + square(d)) / (2 * d)
	y := (square(p1.R)-square(p3.R)+square(i)+square(j))/(2*j) - (i/j)*x

	p4 := add(p1, add(multiply(ex, x), multiply(ey, y)))
	return p4, nil
}
