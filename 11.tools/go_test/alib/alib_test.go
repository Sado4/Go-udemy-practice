package alib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップします")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("3じゃない", v)
		t.Errorf("%d != %d", v, 3)
	}
}
