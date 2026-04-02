package permiter

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
	return 0.0
}

func (c Circle) Area() float64 {
	return 0.0
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
