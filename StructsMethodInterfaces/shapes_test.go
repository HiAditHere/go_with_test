package structsmethodinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	/*checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})*/

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314.1592653589793},
		{Triangle{10, 5}, 25},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("want %g got %g", got, tt.want)
		}
	}
}
