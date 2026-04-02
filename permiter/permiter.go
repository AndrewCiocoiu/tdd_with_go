package permiter

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
