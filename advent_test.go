package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay5(t *testing.T) {
	input := "0 3 0 1 -3"
	out := advent5A(input)
	expected := 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "0 3 0 1 -3"
	out = advent5B(input)
	expected = 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay6(t *testing.T) {
	input := "0 2 7 0"
	out := advent6A(input)
	expected := 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = "0 2 7 0"
	out = advent6B(input)
	expected = 4
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay7(t *testing.T) {
	input := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`
	out := advent7A(input)
	expected := "tknk"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	// intOut := advent7B(input)
	// intExp := 60
	// if !cmp.Equal(intOut, intExp) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(intOut, intExp))
	// }
}

func TestDay8(t *testing.T) {
	input := `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
	out := advent8A(input)
	expected := 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	out = advent8B(input)
	expected = 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay9(t *testing.T) {
	input := "{}"
	out := advent9A(input)
	expected := 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{{}}}"
	out = advent9A(input)
	expected = 6
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{},{}}"
	out = advent9A(input)
	expected = 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{{},{},{{}}}}"
	out = advent9A(input)
	expected = 16
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{<a>,<a>,<a>,<a>}"
	out = advent9A(input)
	expected = 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<ab>},{<ab>},{<ab>},{<ab>}}"
	out = advent9A(input)
	expected = 9
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	out = advent9A(input)
	expected = 9
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	out = advent9A(input)
	expected = 3
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay10(t *testing.T) {
	input := "3,4,1,5"
	iout := advent10A(input)
	iexpected := 12
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}

	input = ""
	out := advent10B(input)
	expected := "a2582a3a0e66e6e86e3812dcb672a272"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "AoC 2017"
	out = advent10B(input)
	expected = "33efeb34ea91902bb2f59c9920caa6cd"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "1,2,3"
	out = advent10B(input)
	expected = "3efbe78a8d82f29979031a4aa0b16a9d"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "1,2,4"
	out = advent10B(input)
	expected = "63960835bcdc130f0b66d7ff4f6a5a8e"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestAdvent(t *testing.T) {
}
