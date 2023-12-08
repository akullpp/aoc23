package main

import "testing"

func TestExampleD3P1(t *testing.T) {
	expectation := 4361

	result := d3p1("input/d3.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD3P1(t *testing.T) {
	expectation := 544664

	result := d3p1("input/d3.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}
