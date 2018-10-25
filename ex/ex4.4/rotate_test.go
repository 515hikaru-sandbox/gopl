package main

import "testing"

func TestRotate(t *testing.T) {
	a := []int{1, 2, 3}
	rotate(a, 2)
	if a[0] != 3 || a[1] != 1 || a[2] != 2 {
		t.Errorf("expected [3 1 2], but got=%v\n", a)
	}
}
