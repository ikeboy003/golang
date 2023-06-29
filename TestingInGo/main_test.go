package main

import "testing"

func TestAdd(t *testing.T) {

	//arrange
	l, r, expect := 1, 2, 4

	val := Add(l, r)
	if expect != val {
		t.Errorf("Failed to add %v and %v. Got %v, Expected %v\n", l, r, expect, val)

	}

}
