package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10}
	got := rec.Perimeter()
	want := 40.0
	outputMessageFloat(t, got, want)
}

func TestArea(t *testing.T) {
	testCases := []struct {
		desc        string
		shape       Shape
		desiredArea float64
	}{
		{"rectangle area", Rectangle{12, 6}, 72.0},
		{"circle area", Circle{10}, 314.1592653589793},
		{"triangle area", Triangle{5, 10}, 25.0},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			shapeArea := tC.shape.Area()
			if shapeArea != tC.desiredArea {
				t.Errorf("%#v got %g want %g", tC.shape, shapeArea, tC.desiredArea)
			}
		})
	}
}

func outputMessageFloat(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
