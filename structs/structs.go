package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
    Base   float64
    Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
    Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
    return 0
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
