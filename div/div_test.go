package sum

import "testing"

func TestDiv(t *testing.T) {
	total := Div(10, 5)
	if total != 2 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMultiply(t *testing.T) {
	total := Multiply(10, 5)
	if total != 50 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 50)
	}
}
