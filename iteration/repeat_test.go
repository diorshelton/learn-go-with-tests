package iteration

import "testing"

func TestRepeat(t *testing.T) {
	output := Repeat("c")
	repeated := "ccccc"

	if output != repeated {
		t.Errorf("expected '%q' but got '%q'", repeated, output)
	}
}
