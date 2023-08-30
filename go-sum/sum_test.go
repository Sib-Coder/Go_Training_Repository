package sum

import (
	"testing"
)

func TestAdder(t *testing.T) {
	a := NewAdder()
	a.Add(1)
	a.Add(1)
	a.Add(1)
	if a.Sum() != 3 {
		t.Errorf("err: want=%v got=%v", 3, a.Sum())
	}
}
