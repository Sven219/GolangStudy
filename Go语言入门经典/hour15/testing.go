package hour15

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatal("The worlf is crumbling")
	}
}
