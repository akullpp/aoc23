package main

import "testing"

func TestExampleD2P1(t *testing.T) {
	expectation := 8

	result := d2p1("input/d2.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD2P1(t *testing.T) {
	expectation := 2685

	result := d2p1("input/d2.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestExampleD2P2(t *testing.T) {
	expectation := 2286

	result := d2p2("input/d2.example")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}

func TestInputD2P2(t *testing.T) {
	expectation := 83707

	result := d2p2("input/d2.input")

	if result != expectation {
		t.Errorf("Expected %d, got %d", expectation, result)
	}
}
