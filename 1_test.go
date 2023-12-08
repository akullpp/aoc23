package main

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

	result := d1p1("input/d1.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestExampleD1P2(t *testing.T) {
	expectation := 281

	result := d1p2("input/d1p2.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD1P2(t *testing.T) {
	expectation := 54728

	result := d1p2("input/d1.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}
