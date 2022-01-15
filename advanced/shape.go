package advanced

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type Triangle struct {
	Side   float64
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Side + t.Base + t.Height
}
