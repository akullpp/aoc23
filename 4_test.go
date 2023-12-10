package main

import "testing"

func TestExampleD4P1(t *testing.T) {
	expectation := 13

	result := d4p1("input/d4.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD4P1(t *testing.T) {
	expectation := 24542

	result := d4p1("input/d4.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}
