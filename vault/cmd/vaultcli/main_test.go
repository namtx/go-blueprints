package main

import "testing"

func TestPop(t *testing.T) {
	args := []string{"one", "two", "three"}
	var s string

	s, args = pop(args)

	if s != "one" {
		t.Errorf("expected 'one', got '%s'", s)
	}

	s, args = pop(args)
	if s != "two" {
		t.Errorf("expected 'two', got '%s'", s)
	}
	s, args = pop(args)
	if s != "three" {
		t.Errorf("expected 'three', got '%s'", s)
	}

	s, args = pop(args)
	if s != "" {
		t.Errorf("expected empty string, got '%s'", s)
	}
}
