package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

const pi = math.Pi

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (circle Circle) Perimeter() float64 {
	return 2 * pi * circle.Radius
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return pi * math.Pow(circle.Radius, 2)
}

func (triangle Triangle) Area() float64 {
	return (triangle.Base * triangle.Height) * 0.5
}

func (triangle Triangle) hypotenuse() float64 {
	return math.Sqrt(math.Pow(triangle.Base, 2) + math.Pow(triangle.Height, 2))
}

func (triangle Triangle) Perimeter() float64 {
	return triangle.Base + triangle.Height + triangle.hypotenuse()
}
