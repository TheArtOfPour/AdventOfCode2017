package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdvent(t *testing.T) {
	input := 6
	out := advent3B(input)
	expected := 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = 10
	out = advent3B(input)
	expected = 11
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = 23
	out = advent3B(input)
	expected = 25
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = 750
	out = advent3B(input)
	expected = 806
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = 325489
	out = advent3B(input)
	fmt.Printf("Result %d\n", out)
}
