package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat string character", func(t *testing.T) {
		output := Repeat("c", 5)
		repeated := "ccccc"
		assertCorrectMessage(t, output, repeated)
	})

	t.Run("specify number of times to repeat char", func(t *testing.T) {
		output := Repeat("r", 3)
		repeated := "rrr"
		assertCorrectMessage(t, output, repeated)
	})
	t.Run("returns 0 when zero entered", func(t *testing.T) {
		output := Repeat("r", 0)
		repeated := ""
		assertCorrectMessage(t, output, repeated)
	})
	t.Run("returns three j's", func(t *testing.T) {
		output := Repeat("j", 3)
		repeated := "jjj"
		assertCorrectMessage(t, output, repeated)
	})
}

func assertCorrectMessage(t testing.TB, output, repeated string) {
	t.Helper()
	if output != repeated {
		t.Errorf("Expected output was %q repeated %q", output, repeated)
	}
}

func ExampleRepeat() {
	repeat:= Repeat("j", 3)
	fmt.Println(repeat)
	//output: jjj
}