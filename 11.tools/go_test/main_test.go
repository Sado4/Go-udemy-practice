package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	if Debug {
		t.Skip("スキップする")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%d != %d", i, 1)
	}
}
