package advanced

import (
	"testing"
)

func TestArea(t *testing.T) {

	areas := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areas {
		got := tt.shape.Area()
		if got != tt.expect {
			t.Errorf("got %.2f, expected %.2f", got, tt.expect)
		}
	}

}

func TestPerimeter(t *testing.T) {
	areas := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{12, 6}, 36},
		{Circle{36}, 226.194},
	}

	for _, tt := range areas {
		got := tt.shape.Perimeter()

		if got != tt.expect {
			t.Errorf("got %.2f, expected %.2f", got, tt.expect)
		}
	}
}
