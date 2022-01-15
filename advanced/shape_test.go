package advanced

import "testing"

type Testable struct {
	name    string
	subject Shape
	expect  float64
}

func TestArea(t *testing.T) {
	areas := []Testable{
		{subject: Rectangle{Height: 50, Width: 10}, expect: 500, name: "Rectangle"},
		{subject: Triangle{Side: 10, Base: 30, Height: 60}, expect: 900, name: "Triangle"},
	}

	for _, tt := range areas {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject.Area()
			if got != tt.expect {
				t.Errorf("%v expected %g but got %g", tt.name, tt.expect, got)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	permiters := []Testable{
		{subject: Rectangle{Height: 50, Width: 10}, expect: 120, name: "Rectangle"},
		{subject: Triangle{Side: 60, Base: 80, Height: 120}, expect: 260, name: "Triangle"},
	}

	for _, tt := range permiters {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject.Perimeter()
			if got != tt.expect {
				t.Errorf("%v expected %g but got %g", tt.name, tt.expect, got)
			}
		})
	}
}
