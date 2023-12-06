package d1

import "testing"

func TestExampleD1P1(t *testing.T) {
	expectation := 142

	result := d1p1("input/d1p1.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD1P1(t *testing.T) {
	expectation := 54916

	result := d1p1("input/d1p1.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}
