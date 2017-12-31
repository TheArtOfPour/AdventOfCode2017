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

func TestDay11(t *testing.T) {
	input := "ne,ne,ne"
	iout := advent11A(input)
	iexpected := 3
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	input = "ne,ne,sw,sw"
	iout = advent11A(input)
	iexpected = 0
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	input = "ne,ne,s,s"
	iout = advent11A(input)
	iexpected = 2
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	input = "se,sw,se,sw,sw"
	iout = advent11A(input)
	iexpected = 3
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay12(t *testing.T) {
	input := `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`
	iout := advent12A(input)
	iexpected := 6
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	iout = advent12B(input)
	iexpected = 2
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay13(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`
	iout := advent13A(input)
	iexpected := 24
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	iout = advent13B(input)
	iexpected = 10
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay14(t *testing.T) {
	input := `flqrgnkx`
	iout := advent14A(input)
	iexpected := 8108
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	iout = advent14B(input)
	iexpected = 1242
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay15(t *testing.T) {
	input := `65,8921`
	iout := advent15A(input)
	iexpected := 588
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	iout = advent15B(input)
	iexpected = 309
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay16(t *testing.T) {
	input := `s1,x3/4,pe/b`
	iout := advent16A(input, "abcde")
	iexpected := "baedc"
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
	iout = advent16B(input, "abcde")
	iexpected = `abcde`
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay17(t *testing.T) {
	input := 3
	iout := advent17A(input)
	iexpected := 638
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay18(t *testing.T) {
	input := `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`
	iout := advent18A(input)
	iexpected := 4
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}
func TestDay19(t *testing.T) {
	input := `    |          
    |  +--+    
    A  |  C    
F---|----E|--+ 
    |  |  |  D 
    +B-+  +--+`
	out := advent19A(input)
	expected := "ABCDEF"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	iout := advent19B(input)
	iexpected := 38
	if !cmp.Equal(iout, iexpected) {
		t.Errorf("Didn't match %s", cmp.Diff(iout, iexpected))
	}
}

func TestDay20(t *testing.T) {
	input := `p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>`
	out := advent20A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay21(t *testing.T) {
	input := `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`
	out := advent21A(input)
	expected := 12
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}
func TestDay22(t *testing.T) {
	input := `..#
#..
...`
	// out := advent22A(input)
	// expected := 5587
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
	out := advent22B(input)
	expected := 2511944
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay24(t *testing.T) {
	input := `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`
	out := advent24A(input)
	expected := 31
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out = advent24B(input)
	expected = 19
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay25(t *testing.T) {
	input := `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
	If the current value is 0:
	- Write the value 1.
	- Move one slot to the right.
	- Continue with state B.
	If the current value is 1:
	- Write the value 0.
	- Move one slot to the left.
	- Continue with state B.

In state B:
	If the current value is 0:
	- Write the value 1.
	- Move one slot to the left.
	- Continue with state A.
	If the current value is 1:
	- Write the value 1.
	- Move one slot to the right.
	- Continue with state A.`
	out := advent25A(input)
	expected := 3
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent25B(input)
	// expected = 19
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestAdvent(t *testing.T) {
}
