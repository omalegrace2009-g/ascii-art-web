package main

import "testing"

func TestProcessText(t *testing.T) {
	input := "add 1E (hex) files, make them nice (cap) and 101 (bin) items (up, 2)"
	expected := "add 30 files, make them Nice and 5 ITEMS"

	result := processText(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}