package utilities

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum != 5 {
		t.Errorf("error")
	}
}
